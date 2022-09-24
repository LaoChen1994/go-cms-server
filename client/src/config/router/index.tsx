import { RouteObject } from "react-router-dom";
import AppLayout from "Components/AppLayout";
import NotFount from "Pages/common/NotFount";
import article from './article'
import overview from "./overview"
import tag from './tag'

const routes: RouteObject = {
  path: "/",
  element: <AppLayout />,
  children: [
    ...overview,
    ...article,
    ...tag,
    {
      path: "/*",
      element: <NotFount />,
    },
  ],
}

export default routes;
