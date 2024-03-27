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
      label={formatMessage({ id: 'component.upstream.fields.checks.passive.healthy.successes' })}
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.passive.healthy.successes.tooltip',
      })}
      required
    >
      <Form.Item
        name={['checks', 'passive', 'healthy', 'successes']}
        noStyle
        initialValue={5}
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'component.upstream.fields.checks.passive.healthy.successes.required',
            }),
          },
        ]}
      >
        <InputNumber disabled={readonly} min={1} max={254} />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
