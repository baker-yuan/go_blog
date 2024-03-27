import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import TimeUnit from './TimeUnit';

const Timeout: React.FC<{
  label: string;
  desc: string;
  name: string[];
  readonly?: boolean;
}> = ({ label, desc, name, readonly }) => {
  const { formatMessage } = useIntl();
  return (
    <Form.Item label={label} required tooltip={desc}>
      <Form.Item
        name={name}
        noStyle
        rules={[
          {
            required: true,
            message: formatMessage({ id: `page.upstream.step.input.${name[1]}.timeout` }),
          },
        ]}
        initialValue={6}
      >
        <InputNumber disabled={readonly} />
      </Form.Item>
      <TimeUnit />
    </Form.Item>
  );
};

export default Timeout;
