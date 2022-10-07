import { Content } from 'Components/AppLayout/AppContent'

import {
  Form, Input, Button, Notification,
} from '@arco-design/web-react'
import useRequest from "Hooks/useRequest";
import { useAuth } from "Context/auth";

const FormItem = Form.Item

interface ILoginForm {
  account: string;
  password: string
}

export default () => {
  const [form] = Form.useForm<ILoginForm>()
  const { loading, request } = useRequest("/api/open/user/login", {
    method: "post",
  })

  const { syncAuth } = useAuth()

  const handleSubmit = async (v: ILoginForm) => {
    try {
      await request(v);
      await syncAuth()
    } catch (e) {
      Notification.error({ content: e instanceof Error ? e.message : "用户登录异常" })
    }
  }

  return (
    <Content style={{ display: "flex", alignItems: "center", justifyContent: "center" }}>
      <Form form={form} layout="horizontal" onSubmit={handleSubmit} wrapperCol={{ span: 6 }} autoComplete="off">
        <FormItem label="账号" field="account" colon rules={[{ required: true }]}>
          <Input name="account" placeholder="请输入账户" />
        </FormItem>
        <FormItem
          label="密码"
          field="password"
          colon
          wrapperCol={{ span: 6 }}
          rules={[{ required: true, minLength: 6, maxLength: 30 }]}
        >
          <Input name="password" type="password" maxLength="30" placeholder="请输入密码" />
        </FormItem>
        <FormItem wrapperCol={{ span: 6, offset: 5 }}>
          <Button htmlType="submit" type="primary" loading={loading}>登录</Button>
          <Button
            htmlType="reset"
            type="outline"
            style={{ marginLeft: "16px" }}
            onClick={() => {
              form.resetFields()
            }}
          >
            重置
          </Button>
        </FormItem>
      </Form>
    </Content>
  )
}
