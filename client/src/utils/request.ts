import axios, { AxiosRequestConfig } from "axios"
import set from "lodash/set"
import { Notification } from '@arco-design/web-react'

import { parse, format } from "url"

async function request<T = any>(config: AxiosRequestConfig): Promise<T> {
  const {
    headers = {}, url, params, data,
  } = config

  if (!url) {
    throw Error("请求地址不能为空")
  }

  if (!headers["Content-Type"]) {
    set(config, ["headers", "Content-Type"], "application/json")
  }

  if (!config.method) {
    config.method = "GET"
  }

  config.withCredentials = true
  config.timeout = 3000

  const { query, ...rest } = parse(url, true)

  if (config.method!.toLowerCase() === "get") {
    config.url = format({
      ...rest,
      query: {
        ...query,
        ...params,
        ...data,
      },
    })
  }

  const res = await axios(config)

  if (res.status > 200) {
    Notification.error(res.data.message)
    throw Error(res.data.message || "接口请求异常")
  }

  return (res.data && res.data.data) || res.data || res || null
}

export default request
