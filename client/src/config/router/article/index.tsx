import { RouteObject } from 'react-router-dom'
import ArticleList from 'Pages/article/list'
import ArticleEdit from 'Pages/article/edit'

const routes: RouteObject[] = [
  {
    path: "/articles",
    element: <ArticleList />,
  },
  {
    path: "/article",
    element: <ArticleEdit />,
    children: [
      {
        path: "/article/:id",
        element: <ArticleEdit />,
      },
    ],
  },
]

export default routes
