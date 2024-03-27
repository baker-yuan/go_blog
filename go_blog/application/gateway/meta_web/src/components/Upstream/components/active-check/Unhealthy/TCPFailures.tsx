import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const TCPFailures: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.Item
      label={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.tcp_failures',
      })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.tcp_failures.tooltip',
      })}
    >
      <Form.Item
        name={['checks', 'active', 'unhealthy', 'tcp_failures']}
        noStyle
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'component.upstream.fields.checks.active.unhealthy.tcp_failures.required',
            }),
          },
        ]}
        initialValue={2}
      >
        <InputNumber disabled={readonly} min={1} max={254} />
      </Form.Item>
    </Form.Item>
  );
};

export default TCPFailures;
