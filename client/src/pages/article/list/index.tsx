import { Content } from 'Components/AppLayout/AppContent'
import useRequest from 'Hooks/useRequest'
import { IArticleItem, IArticlesReq } from 'Types/api'

import { useCallback, useEffect } from "react";
import {
  Button, Space, Table, TableColumnProps, Tag,
} from '@arco-design/web-react'
import { useNavigate } from "react-router-dom";

const columns: TableColumnProps<IArticleItem>[] = [
  {
    title: "文章id",
    dataIndex: "id",
  },
  {
    title: "文章标题",
    dataIndex: "title",
  },
  {
    title: "状态",
    dataIndex: "state",
    render(col) {
      let color = "gray"
      let text = "未知"

      switch (col) {
        case 1:
          color = "green";
          text = "上架"
          break;
        case 2:
          color = "red";
          text = "下架";
          break
        default:
          break;
      }

      return <Tag color={color}>{text}</Tag>
    },
  },
  {
    title: "操作",
    dataIndex: "_",
    render(_, item) {
      const { state } = item;

      return (
        <Button.Group>
          <Button type="text">编辑</Button>
          {state !== 2
                && <Button type="text" disabled={state === 2}>{state === 1 ? "下架" : "上架"}</Button>}
          <Button type="text" status="danger">删除</Button>
        </Button.Group>
      )
    },
  },
]

export default function () {
  const { data, loading, request } = useRequest<IArticlesReq, IArticleItem[]>("/api/cms/articles", {
    data: {
      tag: true,
    },
  })

  const navigate = useNavigate();

  const fetchData = useCallback(async (page = 1) => {
    request({
      page,
      size: 10,
    })
  }, []);

  useEffect(() => {
    fetchData()
  }, [])

  const handleCreate = useCallback(() => {
    navigate('/article')
  }, [])

  return (
    <Content>
      <Space direction="vertical">
        <Space>
          <Button onClick={handleCreate}>新建</Button>
        </Space>
        <Space style={{ width: "100%" }}>
          <Table style={{ width: "100%" }} loading={loading} data={data || []} columns={columns} />
        </Space>
      </Space>
    </Content>
  )
}
