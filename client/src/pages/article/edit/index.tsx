import { Content } from "Components/AppLayout/AppContent";
import {
  Form, Input, FormInstance, Button,
} from '@arco-design/web-react'
import { useCallback, useRef } from "react";
import Editor from "Components/Editor";

const FormItem = Form.Item

const INIT_STATE = {
  title: '',
  summary: '',
  content: '',
}

export default function () {
  const form = useRef<FormInstance | null>(null);

  const handleReset = useCallback(() => {
    form.current?.resetFields()
  }, [])

  const handleSubmit = useCallback(() => {
    const data = form.current?.getFields();
    console.log(data)
  }, [])

  return (
    <Content>
      <Form autoComplete="off" ref={form} initialValues={INIT_STATE}>
        <FormItem label="标题" field="title" rules={[{ required: true, maxLength: 50 }]}>
          <Input />
        </FormItem>
        <FormItem label="简介" field="summary" rules={[{ required: true }]}>
          <Input.TextArea />
        </FormItem>
        <FormItem label="文章内容" field="content" rules={[{ required: true }]}>
          <Editor height={500} />
        </FormItem>
        <FormItem wrapperCol={{ offset: 5 }}>
          <Button type="primary" style={{ marginRight: 24 }} onClick={handleSubmit}>提交</Button>
          <Button type="default" onClick={handleReset}>清空</Button>
        </FormItem>
      </Form>
    </Content>
  )
}
