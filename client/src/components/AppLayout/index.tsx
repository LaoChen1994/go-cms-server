import {
  Layout,
} from '@arco-design/web-react'

import React, { useState } from "react";
import AppContent from "Components/AppLayout/AppContent";
import AppMenu from './AppMenu'
import AppHeader from './AppHeader'

import Styles from './index.module.scss'

export default function () {
  const [collasped, setCollasped] = useState(false)

  const handleCollasped = React.useCallback(() => {
    setCollasped(!collasped)
  }, [collasped])

  return (
    <Layout>
      <Layout.Sider
        className={Styles.sideBar}
        breakpoint="lg"
        theme="light"
        collapsible
        collapsed={collasped}
        onCollapse={handleCollasped}
        collapsedWidth={40}
      >
        <AppMenu isCollapse={collasped} />
      </Layout.Sider>
      <Layout className={Styles.content}>
        <Layout.Header>
          <AppHeader />
        </Layout.Header>
        <Layout.Content>
          <AppContent />
        </Layout.Content>
      </Layout>
    </Layout>
  )
}
