import { useOutlet } from 'react-router-dom'
import { PropsWithChildren, useEffect, useMemo } from "react";
import cx from 'classnames'
import { useAuth } from "Context/auth";
import useBreadcrumb from "Hooks/useBreadcrumb";
import Styles from './index.module.scss';

export function Content(props: PropsWithChildren<any>) {
  const { children, ...resProps } = props
  const breadcrumb = useBreadcrumb()
  const { className, style, ...res } = resProps

  const height = useMemo(() => (breadcrumb ? 88 : 32), [breadcrumb])

  return (
    <>
      <div className={Styles.breadWrapper}>{breadcrumb}</div>
      <div
        className={cx(Styles.layoutContent, className)}
        {...res}
        style={{ ...style, height: `calc(100% - ${height}px)` }}
      >
        {children}
      </div>
    </>
  )
}

export default function () {
  const element = useOutlet()
  const { syncAuth, auth } = useAuth()

  useEffect(() => {
    if (!auth.id) {
      syncAuth()
    }
  }, [])

  return element
}
