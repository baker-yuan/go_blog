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
      label={formatMessage({ id: 'component.upstream.fields.checks.active.healthy.successes' })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.healthy.successes.tooltip',
      })}
    >
      <Form.Item
        name={['checks', 'active', 'healthy', 'successes']}
        noStyle
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'component.upstream.fields.checks.active.healthy.successes.required',
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

export default Component;
