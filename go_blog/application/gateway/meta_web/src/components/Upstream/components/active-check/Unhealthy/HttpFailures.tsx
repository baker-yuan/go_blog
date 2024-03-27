import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.Item
      label={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.http_failures',
      })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.http_failures.tooltip',
      })}
    >
      <Form.Item
        name={['checks', 'active', 'unhealthy', 'http_failures']}
        noStyle
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'component.upstream.fields.checks.active.unhealthy.http_failures.required',
            }),
          },
        ]}
        initialValue={5}
      >
        <InputNumber disabled={readonly} min={1} max={254} />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
