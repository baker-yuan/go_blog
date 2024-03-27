import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import TimeUnit from '../../TimeUnit';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();
  return (
    <Form.Item
      label={formatMessage({ id: 'component.upstream.fields.checks.active.healthy.interval' })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.healthy.interval.tooltip',
      })}
    >
      <Form.Item
        noStyle
        style={{ marginBottom: 0 }}
        name={['checks', 'active', 'healthy', 'interval']}
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'page.upstream.step.input.healthyCheck.activeInterval',
            }),
          },
        ]}
        initialValue={1}
      >
        <InputNumber disabled={readonly} min={1} />
      </Form.Item>
      <TimeUnit />
    </Form.Item>
  );
};

export default Component;
