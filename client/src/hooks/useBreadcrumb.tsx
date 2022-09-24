import { useLoaderData } from 'react-router-dom'
import { Breadcrumb } from "@arco-design/web-react";
import { IMyBreadItemConfig } from "Utils/loader";

const BreadItem = Breadcrumb.Item

function useBreadcrumb() {
  const breadCrumb = (useLoaderData() || {}) as { breadItems: IMyBreadItemConfig[] };

  if (!breadCrumb.breadItems) {
    return null
  }

  return (
    <Breadcrumb>
      {breadCrumb.breadItems.map((item) => (
        <BreadItem key={`breadcrumb_${item.title}`}>
          {
                item.href ? <a href={item.href}>{item.title}</a> : item.title
              }
        </BreadItem>
      ))}
    </Breadcrumb>
  )
}

export default useBreadcrumb
