import * as dotenv from "dotenv";
import { Connector } from "./connector";
import { file_pumpfun_pumpfun } from "./pumpfun/pumpfun_pb";
import { Signal } from "./signal";

dotenv.config();

async function main() {
  const connector = new Connector();

  await connector.registerDescriptor(file_pumpfun_pumpfun);

  const signal = new Signal("SIGINT");

  await connector.run(signal);

  console.log("Connector shutdown");

  process.exit(0);
}
main();
