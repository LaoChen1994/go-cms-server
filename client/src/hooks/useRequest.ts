import request from "Utils/request";
import type { AxiosRequestConfig } from "axios";
// import debounce from 'lodash/debounce'
// import throttle from 'lodash/throttle'
import omit from 'lodash/omit'
import { useEffect, useState } from "react";
import { format, parse } from 'url'
import { getConfig } from "Utils/index";

interface IRequestOptios<D> extends Omit<AxiosRequestConfig<D>, "url"> {
  auto?: boolean
  throttle?: boolean
  debounce?: boolean
}

export default function <D, T = any> (url: string, options?: IRequestOptios<D>) {
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")
  const [data, setData] = useState<T>({} as T)
  const { auto = false } = options || {}

  const wrapperRequest = async (args?: D) => {
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
      });

      setData(rlt)
      setLoading(false)

      return rlt
    } catch (e) {
      setError(e instanceof Error ? e.message : "请求异常")
    } finally {
      setLoading(false)
    }
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
