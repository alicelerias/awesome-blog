import Redis from "ioredis";
import { promisify } from "util";

export const redisClient = new Redis();

export const getRedis = (value: string) => {
  const syncRedisGet = promisify(redisClient.get).bind(redisClient);
  return syncRedisGet(value);
};

export const setRedis = (key: string, value: string) => {
  const syncRedisSet = promisify(redisClient.set).bind(redisClient);
  return syncRedisSet(key, value);
};
