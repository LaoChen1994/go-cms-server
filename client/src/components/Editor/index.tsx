import MDEditor, { commands, MDEditorProps } from '@uiw/react-md-editor';
import {
  memo, useState, PropsWithRef, FC, useMemo,
} from "react";
import Styles from './index.module.scss';

interface IEditor extends MDEditorProps {

}

const Editor: FC<PropsWithRef<IEditor>> = memo((props) => {
  const { value, onChange, ...res } = props;
  const [state, setState] = useState<string>();

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

  return (
    <div className={Styles.container}>
      <MDEditor value={memoState} onChange={handleChange} {...res} />
    </div>
  )
})

export default Editor
