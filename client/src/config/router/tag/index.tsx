import type { RouteObject } from 'react-router-dom'
import TagList from 'Pages/tag/list'
import TagEdit from 'Pages/tag/edit'

const routes: RouteObject[] = [
  {
    path: "/tags",
    element: <TagList />,
  },
  {
    path: '/tag',
    element: <TagEdit />,
    children: [
      {
        path: '/tag/:id',
        element: <TagEdit />,
      },
    ],
  },
]

export default routes
