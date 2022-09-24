import type { RouteObject } from 'react-router-dom'
import CMSOverview from 'Pages/overview/cms'
import SMSOverview from 'Pages/overview/sms'

const routes: RouteObject[] = [
  {
    path: "/overview",
    children: [
      {
        path: "/overview/cms",
        element: <CMSOverview />,
      },
      {
        path: '/overview/sms',
        element: <SMSOverview />,
      },
    ],
  },
]

export default routes
