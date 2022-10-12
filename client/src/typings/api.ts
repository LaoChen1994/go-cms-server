export interface ICommonField {
    createdAt: string;
    updatedAt: string;
    deletedAt: string | null;
    createdId: string;
}

export interface IArticlesReq {
    tag?: boolean;
    page?: number;
    size?: number
}

export interface IArticleItem extends ICommonField {
    id: number;
    title: string;
    content: string;
    desc: string;
    state: number;
    createdAt: string;
    updatedAt: string;
}

export interface ITag extends ICommonField {
    id: number;
    name: string;
    state: number;
}
