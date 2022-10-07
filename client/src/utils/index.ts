import config from "Config/static"
import get from 'lodash/get'

export { default as breadLoader } from './loader';
export { default as request } from './request'
export { default as composeLoader } from './compose'

type IStaticConfig = typeof config

export function getConfig<T extends keyof IStaticConfig, K extends keyof IStaticConfig[T]>(key: K) {
  const env = process.env.START_ENV as keyof IStaticConfig

  return get(config, [env, key])
}
