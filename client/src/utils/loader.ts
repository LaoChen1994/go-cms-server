import type { BreadCrumbItemProps } from "@arco-design/web-react"

export interface IMyBreadItemConfig extends BreadCrumbItemProps {
    title?: string;
    href?: string;
}

const breadLoader = (props: IMyBreadItemConfig[]) => ({
  breadItems: props,
})

export default breadLoader
