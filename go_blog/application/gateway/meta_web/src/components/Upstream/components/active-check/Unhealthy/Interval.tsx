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
      label={formatMessage({ id: 'component.upstream.fields.checks.active.unhealthy.interval' })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.interval.tooltip',
      })}
    >
      <Form.Item
        name={['checks', 'active', 'unhealthy', 'interval']}
        noStyle
        initialValue={1}
        rules={[
          {
            required: true,
            message: formatMessage({
              id: 'page.upstream.step.input.healthyCheck.activeInterval',
            }),
          },
        ]}
      >
        <InputNumber disabled={readonly} min={1} />
      </Form.Item>
      <TimeUnit />
    </Form.Item>
  );
};

export default Component;
