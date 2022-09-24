import { RouteObject } from 'react-router-dom'
import ArticleList from 'Pages/article/list'
import ArticleEdit from 'Pages/article/edit'
import breadLoader from "Utils/loader";

const routes: RouteObject[] = [
  {
    path: "/articles",
    element: <ArticleList />,
    loader() {
      return breadLoader([{
        title: "首页",
        href: "/",
      }, {
        title: "文章列表",
      }])
    },
  },
  {
    path: "/article",
    element: <ArticleEdit />,
    children: [
      {
        path: "/article/:id",
        element: <ArticleEdit />,
        loader: ({ params }) => breadLoader([{
          title: "首页",
          href: "/",
        }, {
          title: "文章列表",
          href: "/articles",
        }, {
          title: "文章编辑",
          href: `/article/${params}`,
        }]),
      },
    ],
  },
]

export default routes
