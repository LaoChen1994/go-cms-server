import type { ReactElement } from 'react'

export interface IMenumConfig {
    title: string;
    path?: string;
    key?: string;
    icon?: ReactElement;
    children?: IMenumConfig[]
}
