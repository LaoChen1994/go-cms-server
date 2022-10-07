import { observer } from 'mobx-react'
import { useMemo } from "react";
import { useAuth } from "Context/auth";
import { Popover } from '@arco-design/web-react'
import Styles from './index.module.scss'

function User() {
  const { auth, signOut } = useAuth()
  const { userInfo } = auth;
  const { name, mobile } = userInfo || {}

  const mobileText = useMemo(() => {
    if (!mobile) return ""

    return mobile.replace(/(\d{3})(\d{4})(\d{4})/i, (a, s1, b, s3) => `${s1} **** ${s3}`)
  }, [mobile])

  const renderContent = () => (
    <div className={Styles.menu}>
      <div
        className={Styles.menuItem}
        onClick={() => signOut()}
        role="presentation"
      >
        登出
      </div>
      <div className={Styles.menuItem}>修改信息</div>
    </div>
  )

  const renderUser = () => (
    <Popover title="登录操作" content={renderContent()} position="bl">

      <div className={Styles.title}>
        <span>
          用户:
          {' '}
          {name}
        </span>
        {mobileText && <span>{mobileText}</span>}
      </div>
    </Popover>

  )

  return (
    auth.id ? renderUser() : <div>请登录</div>
  )
}

export default observer(User)
