import { Connection, SolanaJSONRPCError } from "@solana/web3.js";
import AsyncLock from "async-lock";
import { LRUCache } from "lru-cache";

export async function sleep(ms: number) {
  await new Promise((resolve) => setTimeout(resolve, ms));
}

const lock = new AsyncLock();
const cache = new LRUCache<number, number>({ max: 100000 });

export async function getBlockTime(connection: Connection, slot: number) {
  return lock.acquire(slot.toString(), async () => {
    let blockTime = cache.get(slot);

    if (blockTime == undefined) {
      while (true) {
        try {
          blockTime = (await connection.getBlockTime(slot))!;
          cache.set(slot, blockTime);
        } catch (e) {
          if (
            e instanceof SolanaJSONRPCError &&
            e.code == -32004 &&
            e.message.includes("Block not available")
          ) {
            await sleep(100);
            continue;
          }
          throw e;
        }
        break;
      }
    }

    return blockTime;
  });
}
