import type { RouteObject } from "react-router-dom"
import Login from "Pages/login"

const routes: RouteObject[] = [{
  path: "/login",
  element: <Login />,
}]

export default routes
