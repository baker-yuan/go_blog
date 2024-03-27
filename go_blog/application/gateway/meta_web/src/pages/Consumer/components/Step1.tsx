import { Form, Input } from 'antd';
import type { FormInstance } from 'antd/lib/form';
import React from 'react';
import { useIntl } from 'umi';

const FORM_LAYOUT = {
  labelCol: {
    span: 2,
  },
  wrapperCol: {
    span: 8,
  },
};

type Props = {
  form: FormInstance;
  disabled?: boolean;
};

const Step1: React.FC<Props> = ({ form, disabled }) => {
  const { formatMessage } = useIntl();
  return (
    <Form {...FORM_LAYOUT} form={form}>
      <Form.Item
        label={formatMessage({ id: 'page.consumer.username' })}
        name="username"
        help={formatMessage({ id: 'component.global.form.itemExtraMessage.nameGloballyUnique' })}
        rules={[
          { required: true },
          {
            pattern: new RegExp(/^[a-zA-Z0-9_]+$/, 'g'),
            message: formatMessage({ id: 'page.consumer.form.itemRuleMessage.username' }),
          },
        ]}
      >
        <Input
          placeholder={formatMessage({ id: 'page.consumer.username.required' })}
          disabled={disabled || window.location.pathname.indexOf('edit') !== -1}
        />
      </Form.Item>
      <Form.Item label={formatMessage({ id: 'component.global.description' })} name="desc">
        <Input.TextArea
          placeholder={formatMessage({ id: 'component.global.description.required' })}
          disabled={disabled}
        />
      </Form.Item>
    </Form>
  );
};

export default Step1;
