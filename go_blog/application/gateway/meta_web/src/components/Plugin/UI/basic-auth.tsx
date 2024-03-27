import { Form, Input } from 'antd';
import type { FormInstance } from 'antd/es/form';
import React from 'react';

type Props = {
  form: FormInstance;
  schema: Record<string, any> | undefined;
  ref?: any;
};

export const FORM_ITEM_LAYOUT = {
  labelCol: {
    span: 4,
  },
  wrapperCol: {
    span: 8,
  },
};

const BasicAuth: React.FC<Props> = ({ form, schema }) => {
  const required: string[] = schema?.required;
  return (
    <Form form={form} {...FORM_ITEM_LAYOUT}>
      <Form.Item
        label="username"
        name="username"
        rules={[{ required: required.indexOf('username') > -1 }]}
        validateTrigger={['onChange', 'onBlur', 'onClick']}
      >
        <Input></Input>
      </Form.Item>
      <Form.Item
        label="password"
        name="password"
        rules={[{ required: required.indexOf('password') > -1 }]}
        validateTrigger={['onChange', 'onBlur', 'onClick']}
      >
        <Input></Input>
      </Form.Item>
    </Form>
  );
};

export default BasicAuth;
