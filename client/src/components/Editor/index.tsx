import MDEditor, { MDEditorProps } from '@uiw/react-md-editor';
import {
  memo, useState, PropsWithRef, FC, useMemo, useEffect, ReactElement,
} from "react";
import { compile as MDXCompile, run } from '@mdx-js/mdx'
import * as runtime from "react/jsx-runtime"
import { MDXProvider } from "@mdx-js/react";
import Styles from './index.module.scss';

interface IEditor extends MDEditorProps {

}

const str = "# 123\n ## 123 \n [数据data](https://www.baidu.com) \n"

const Editor: FC<PropsWithRef<IEditor>> = memo((props) => {
  const { value, onChange, ...res } = props;
  const [state, setState] = useState<string>();

  const [Component, setComponent] = useState<(() => ReactElement) | null>(null)

  const memoState = useMemo(() => {
    if (value !== undefined) return value;

    return state
  }, [value, state])

  const handleChange: MDEditorProps['onChange'] = (newStr, event, newState) => {
    if (onChange) {
      onChange(newStr, event, newState)
    } else {
      setState(newStr)
    }
  }

  useEffect(() => {
    (async () => {
      const vFile = await MDXCompile(str, { outputFormat: "function-body" }) || {};
      const Content = await run(vFile, runtime)
      setComponent(() => Content.default)
    })()
  }, [])

  return (
    <div className={Styles.container}>
      <MDEditor value={memoState} onChange={handleChange} {...res} />
      {Component ? (
        <MDXProvider>
          <Component />
        </MDXProvider>
      ) : null}
    </div>
  )
})

export default Editor
