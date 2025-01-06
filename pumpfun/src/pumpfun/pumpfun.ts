import { create, toBinary } from "@bufbuild/protobuf";
import { TimestampSchema } from "@bufbuild/protobuf/wkt";
import { BN, BorshCoder, EventParser, Idl } from "@project-serum/anchor";
import { Connection, PublicKey } from "@solana/web3.js";
import Decimal from "decimal.js";
import { getBlockTime } from "../utils";
import IDL from "./pumpfun.idl.json";
import {
  CompleteEvent,
  CompleteEventSchema,
  CreateEvent,
  CreateEventSchema,
  SetParamsEvent,
  SetParamsEventSchema,
  TradeEvent,
  TradeEventSchema,
} from "./pumpfun_pb";

export type EventMessage =
  | CreateEvent
  | TradeEvent
  | CompleteEvent
  | SetParamsEvent;

export const PROGRAM_ID = new PublicKey(
  "6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P"
);

export function encode(msg: EventMessage) {
  switch (msg.$typeName) {
    case "pumpfun.CreateEvent":
      return toBinary(CreateEventSchema, msg);
    case "pumpfun.TradeEvent":
      return toBinary(TradeEventSchema, msg);
    case "pumpfun.CompleteEvent":
      return toBinary(CompleteEventSchema, msg);
    case "pumpfun.SetParamsEvent":
      return toBinary(SetParamsEventSchema, msg);
  }
}

const eventParser = new EventParser(PROGRAM_ID, new BorshCoder(IDL as Idl));

export async function parseLogs(
  connection: Connection,
  logs: string[],
  txHash: string,
  slot: number,
  blockTime?: number | null
) {
  const msgs = [];
  const events = eventParser.parseLogs(logs);

  let logIndex = 0;

  for (const event of events) {
    switch (event.name) {
      case "CreateEvent":
        msgs.push(
          create(CreateEventSchema, {
            ts: create(TimestampSchema, {
              seconds: BigInt(
                blockTime ?? (await getBlockTime(connection, slot))
              ),
            }),
            txHash,
            logIndex,
            name: event.data.name,
            symbol: event.data.symbol,
            uri: event.data.uri,
            mint: event.data.mint.toBase58(),
            bondingCurve: event.data.bondingCurve.toBase58(),
            user: event.data.user.toBase58(),
          })
        );
        break;
      case "TradeEvent":
        msgs.push(
          create(TradeEventSchema, {
            ts: create(TimestampSchema, {
              seconds: BigInt(event.data.timestamp),
            }),
            txHash,
            logIndex,
            mint: event.data.mint.toBase58(),
            solAmount: tokenAmount(event.data.solAmount, 9),
            tokenAmount: tokenAmount(event.data.tokenAmount, 6),
            isBuy: event.data.isBuy,
            user: event.data.user.toBase58(),
            virtualSolReserves: tokenAmount(event.data.virtualSolReserves, 9),
            virtualTokenReserves: tokenAmount(
              event.data.virtualTokenReserves,
              6
            ),
          })
        );
        break;
      case "CompleteEvent":
        msgs.push(
          create(CompleteEventSchema, {
            ts: create(TimestampSchema, {
              seconds: BigInt(event.data.timestamp),
            }),
            txHash,
            logIndex,
            user: event.data.user.toBase58(),
            mint: event.data.mint.toBase58(),
            bondingCurve: event.data.bondingCurve.toBase58(),
          })
        );
        break;
      case "SetParamsEvent":
        msgs.push(
          create(SetParamsEventSchema, {
            ts: create(TimestampSchema, {
              seconds: BigInt(
                blockTime ?? (await getBlockTime(connection, slot))
              ),
            }),
            txHash,
            logIndex,
            feeRecipient: event.data.feeRecipient.toBase58(),
            initialVirtualTokenReserves: tokenAmount(
              event.data.initialVirtualTokenReserves,
              6
            ),
            initialVirtualSolReserves: tokenAmount(
              event.data.initialVirtualSolReserves,
              9
            ),
            initialRealTokenReserves: tokenAmount(
              event.data.initialRealTokenReserves,
              6
            ),
            tokenTotalSupply: tokenAmount(event.data.tokenTotalSupply, 6),
            feeBasisPoints: BigInt(event.data.feeBasisPoints.toString(10)),
          })
        );
        break;
      default:
        console.log(`Invalid event: ${event.name}`);
    }

    logIndex++;
  }

  return msgs;
}

function tokenAmount(amount: BN, decimals: number) {
  return new Decimal(amount.toString(10))
    .div(new Decimal(10).pow(decimals))
    .toFixed();
}
