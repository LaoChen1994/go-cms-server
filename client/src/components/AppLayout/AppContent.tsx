import { useOutlet } from 'react-router-dom'
import Styles from './index.module.scss'

export default function () {
  const element = useOutlet()

  return (
    <div className={Styles.main}>
      {element}
    </div>
  )
}
