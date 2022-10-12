import request from "Utils/request";
import type { AxiosRequestConfig } from "axios";
// import debounce from 'lodash/debounce'
// import throttle from 'lodash/throttle'
import omit from 'lodash/omit'
import { useEffect, useState } from "react";
import { format, parse } from 'url'
import { getConfig } from "Utils/index";

interface IRequestOptios<D, T> extends Omit<AxiosRequestConfig<D>, "url"> {
  auto?: boolean
  throttle?: boolean
  debounce?: boolean
  format?: (data: T) => unknown
}

export default function <D = any, T = any> (
  url: string,
  options?: IRequestOptios<D, T>,
) {
  const { auto = false } = options || {}
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [data, setData] = useState<T | null>(null)

  const wrapperRequest = async (args?: D): Promise<T> => {
    const requestOptions = omit(options, ["auto", "throttle", "debounce"])
    setLoading(true)
    let requestUrl = url;

    if (!url.startsWith("http")) {
      const parseUrl = parse(url);
      requestUrl = format({
        ...parseUrl,
        host: getConfig("ApiHost"),
        protocol: process.env.NODE_ENV === "production" ? "https" : "http",
      })
    }

    try {
      const rlt = await request<T>({
        url: requestUrl,
        ...requestOptions,
        data: {
          ...(requestOptions.data || {}),
          ...(args || {}),
        },
      }) as T;

      setData(rlt)
      setLoading(false)

      return rlt
    } catch (e) {
      setError(e instanceof Error ? e.message : "请求异常")
    } finally {
      setLoading(false)
    }

    return {} as T
  }

  useEffect(() => {
    if (auto) {
      wrapperRequest()
    }
  }, [])

  return {
    loading,
    data,
    error,
    request: wrapperRequest,
  }
}
