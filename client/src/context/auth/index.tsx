import { makeAutoObservable } from 'mobx'
import {
  createContext, PropsWithChildren, useContext, useEffect,
} from 'react'
import { useLocation, useNavigate } from "react-router-dom";
import useRequest from "Hooks/useRequest";
import { IUserInfo } from './interface'

export class Auth {
  userInfo: IUserInfo | null = null;

  constructor() {
    makeAutoObservable(this)
  }

  get name() {
    return this.userInfo?.name || ""
  }

  set name(name: string) {
    if (!this.userInfo) {
      return
    }

    this.userInfo.name = name
  }

  get id() {
    return this.userInfo?.id || 0
  }

  setUserInfo(userInfo: IUserInfo | null) {
    this.userInfo = userInfo
  }
}

const auth = new Auth();

export const AuthContext = createContext<Auth>(auth)

export function AuthProvider(props: PropsWithChildren<any>) {
  const { children } = props
  return (
    <AuthContext.Provider value={auth}>
      {children}
    </AuthContext.Provider>
  )
}

let requestCache: Promise<IUserInfo | undefined> | null = null

export const useAuth = () => {
  const status = useContext(AuthContext)
  const { request } = useRequest<unknown, IUserInfo>("/api/open/user/auth", {
    method: "get",
  })

  const { request: signOut } = useRequest("/api/open/user/loginout", {
    method: "post",
  })

  const location = useLocation()
  const navigate = useNavigate()

  const syncAuth = async () => {
    const isLoginPage = location.pathname.includes("login")
    let loading = false
    if (status.id && isLoginPage) {
      navigate("/")
      return
    }

    if (!requestCache) {
      requestCache = request()
      loading = true
    }

    const userInfo = await requestCache

    if (loading) {
      requestCache = null
    }

    if (userInfo) {
      status.setUserInfo(userInfo)
      if (isLoginPage) {
        navigate("/")
      }
      return
    }

    if (isLoginPage) {
      return
    }

    status.setUserInfo(null)
    navigate("/login")
  }

  const handleSignout = async () => {
    await signOut();
    await syncAuth();
    const isLoginPage = location.pathname.includes("login")

    if (!isLoginPage) {
      navigate("/login")
    }
  }

  useEffect(() => {
    if (!auth.id) {
      syncAuth()
    }
  }, [auth.id])

  return {
    syncAuth,
    auth: status,
    signOut: handleSignout,
  }
}

export default auth
