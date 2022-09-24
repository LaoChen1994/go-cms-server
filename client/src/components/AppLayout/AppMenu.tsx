import { Menu, MenuProps } from '@arco-design/web-react'
import { useNavigate } from 'react-router-dom'
import React, { useEffect, useState } from "react";

import menu from 'Config/menu'
import { IMenumConfig } from "Types/route"
import { MENU_TITLE, COLLASPE_MENU_TITLE } from "Constant/layout"
import cx from 'classnames'
import Style from './index.module.scss'

interface IMenuProps {
  isCollapse: boolean
}

const routeCache: Record<string, string> = {}

const renderMenuItem = (item: IMenumConfig, index: number, prefix: string = '') => {
  const {
    icon, title, path = "/", children, key,
  } = item
  const mergeKey = key || `${prefix}_${index}`
  const hasSubMenu = !!children?.length

  const mergeTitle = icon ? (
    <>
      {icon}
      {title}
    </>
  ) : title

  if (!hasSubMenu) {
    routeCache[mergeKey] = path!
    return <Menu.Item key={mergeKey}>{mergeTitle}</Menu.Item>
  }

  return (
    <Menu.SubMenu key={mergeKey} title={mergeTitle}>
      {children!.map((subItem, i) => renderMenuItem(subItem, i, mergeKey))}
    </Menu.SubMenu>
  )
}

const AppMenu: React.FC<IMenuProps> = (props) => {
  const { isCollapse } = props
  const navigate = useNavigate()
  const [title, setTitle] = useState(MENU_TITLE)

  const handleSubItemClick: MenuProps["onClickMenuItem"] = (key) => {
    const path = routeCache[key];
    navigate(path)
  }

  useEffect(() => {
    setTitle(isCollapse ? COLLASPE_MENU_TITLE : MENU_TITLE)
  }, [isCollapse])

  return (
    <>
      <div className={cx(Style.title, isCollapse ? Style.toggle : "")}>
        {title}
      </div>
      <Menu
        onClickMenuItem={handleSubItemClick}
        theme="light"
        className={Style.menu}
        collapse={isCollapse}
      >
        {menu.map((item, i) => renderMenuItem(item, i))}
      </Menu>
    </>
  )
}

export default AppMenu
