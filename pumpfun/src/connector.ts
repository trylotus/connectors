import { DescFile, toBinary } from "@bufbuild/protobuf";
import { FileDescriptorProtoSchema } from "@bufbuild/protobuf/wkt";
import { Connection } from "@solana/web3.js";
import axios from "axios";
import * as fs from "fs";
import * as Joi from "joi";
import * as yaml from "js-yaml";
import { Kafka, KafkaConfig, Producer, ProducerRecord } from "kafkajs";
import { Queue } from "queue-typescript";
import { Database } from "./database";
import { EventMessage, PROGRAM_ID, encode, parseLogs } from "./pumpfun/pumpfun";
import { Signal } from "./signal";
import { sleep } from "./utils";

const schema = Joi.object({
  DB_URL: Joi.string()
    .required()
    .uri({ scheme: ["postgresql"] }),
  RPC_URL: Joi.string()
    .required()
    .uri({ scheme: ["https"] }),
  REGISTRY_URL: Joi.string()
    .required()
    .uri({ scheme: ["http", "https"] }),
  KAFKA_CONFIG: Joi.string(),
});

export class Connector {
  private readonly name: string;
  private readonly author: string;
  private readonly version: string;

  private readonly registryUrl: string;

  private readonly database: Database;
  private readonly connection: Connection;
  private readonly producer: Producer;

  constructor() {
    const { name, author, version } = yaml.load(
      fs.readFileSync("manifest.yaml", "utf8")
    ) as any;

    this.name = name;
    this.author = author;
    this.version = version;

    const { value, error } = schema.validate(process.env, {
      allowUnknown: true,
    });

    if (error) throw error;

    this.registryUrl = value.REGISTRY_URL;
    this.database = new Database(value.DB_URL);
    this.connection = new Connection(value.RPC_URL, "confirmed");

    const kafkaConfig: KafkaConfig = JSON.parse(
      fs.readFileSync(value.KAFKA_CONFIG ?? "kafka.json", "utf8")
    );

    this.producer = new Kafka(kafkaConfig).producer();
  }

  get id() {
    return `${this.author}.${this.name}.${this.version.replace(/\./g, "_")}`;
  }

  async registerDescriptor(...files: DescFile[]) {
    await axios.put(`${this.registryUrl}/api/v1/descriptors`, {
      author: this.author,
      connector: this.name,
      version: this.version,
      files: files.map((file) =>
        Buffer.from(toBinary(FileDescriptorProtoSchema, file.proto)).toString(
          "base64"
        )
      ),
    });
  }

  async run(signal: Signal) {
    await Promise.all([this.producer.connect(), this.database.initialize()]);

    const { latestSignature } = await this.database.getProgram(
      PROGRAM_ID.toBase58()
    );

    await Promise.all([
      this.subscribe(signal, latestSignature),
      // this.backfill(signal, earliestSignature),
    ]);
  }

  private async subscribe(signal: Signal, after?: string) {
    interface EventLogs {
      logs: string[];
      signature: string;
      slot: number;
    }

    const queue = new Queue<EventLogs>();

    const subscriptionId = this.connection.onLogs(
      PROGRAM_ID,
      ({ logs, err, signature }, { slot }) => {
        if (err) return;

        queue.enqueue({ logs, signature, slot });
      },
      "confirmed"
    );

    while (!signal.done) {
      if (queue.length == 0) {
        await sleep(100);
        continue;
      }

      const { logs, signature, slot } = queue.dequeue();

      if (after) {
        await this.backfill(signal, signature, after); // TODO: send to fct channel
        after = undefined;
      }

      const msgs = await parseLogs(this.connection, logs, signature, slot);

      await this.produce("fct", msgs);

      await this.database.setLatestSignature(PROGRAM_ID.toBase58(), signature);
    }

    await this.connection.removeOnLogsListener(subscriptionId);
  }

  private async backfill(signal: Signal, before?: string, until?: string) {
    while (!signal.done) {
      const signatures = (
        await this.connection.getSignaturesForAddress(
          PROGRAM_ID,
          { before, until },
          "confirmed"
        )
      ).map((info) => info.signature);

      if (signatures.length == 0) break;

      for (const signature of signatures) {
        if (signal.done) return;

        const tx = await this.connection.getTransaction(signature, {
          maxSupportedTransactionVersion: 0,
        });

        if (!tx || tx.meta?.err) continue;

        const msgs = await parseLogs(
          this.connection,
          tx.meta?.logMessages ?? [],
          signature,
          tx.slot,
          tx.blockTime
        );

        await this.produce("bf", msgs);
      }

      before = signatures[signatures.length - 1];

      await this.database.setEarliestSignature(PROGRAM_ID.toBase58(), before);
    }
  }

  private async produce(msgType: "bf" | "fct", msgs: EventMessage[]) {
    const topicMessages: ProducerRecord[] = msgs.map((msg) => {
      return {
        topic: `${msgType}.${this.id}.${msg.$typeName.replace(".", "_")}`,
        messages: [
          {
            key: `${msg.txHash}|${msg.logIndex}`,
            value: Buffer.from(encode(msg)),
          },
        ],
      };
    });

    const results = await this.producer.sendBatch({ topicMessages });

    for (const { topicName: topic, errorCode } of results) {
      if (errorCode)
        console.error("Delivery failed", {
          topic,
          errorCode,
        });
      else
        console.log("Delivered message", {
          topic,
        });
    }
  }
}
