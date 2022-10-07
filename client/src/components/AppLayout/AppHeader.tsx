import User from "Components/User"
import Styles from './index.module.scss';

function Header() {
  return (
    <div className={Styles.header}>
      <User />
    </div>
  )
}

export default Header
