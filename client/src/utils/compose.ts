import { LoaderFunction } from "react-router-dom";

export default function composeLoader(...fns: LoaderFunction[]): LoaderFunction {
  return (...args) => fns.reduce(async (prev, fn) => {
    let rlt = fn(...args)

    if (rlt instanceof Promise) {
      rlt = await rlt
    }

    return ({ ...prev, ...rlt })
  }, {})
}
