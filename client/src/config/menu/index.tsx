import type { IMenumConfig } from 'Types/route'
import {
  IconHome, IconCode, IconTag, IconFile, IconDashboard,
} from '@arco-design/web-react/icon'

const menu: IMenumConfig[] = [
  {
    title: "工作台",
    path: "/overview",
    icon: <IconHome />,
    children: [
      {
        path: "/overview/cms",
        icon: <IconDashboard />,
        title: "工作台",
      },
    ],
  },
  {
    title: "文章",
    icon: <IconFile />,
    children: [
      {
        path: "/articles",
        title: "文章列表",
        icon: <IconCode />,
      },
      {
        path: "/tags",
        title: "标签",
        icon: <IconTag />,
      },
    ],
  },
]

export default menu
