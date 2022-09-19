import ArticleList from 'Pages/article/list'
import ArticleEdit from 'Pages/article/edit'
import { IGetRoutes } from "Types/route";

const getConfig: IGetRoutes = () => ({
  routes: [
    {
      path: "/articles",
      element: <ArticleList />,
    },
    {
      path: "/article/:id",
      element: <ArticleEdit />,
    },
    {
      path: "/article",
      element: <ArticleEdit />,
    },
  ],
})

export default getConfig
