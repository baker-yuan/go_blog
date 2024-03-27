import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import TimeUnit from '../TimeUnit';

type Props = {
  readonly?: boolean;
};

const ActiveCheckTimeoutComponent: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.Item
      label={formatMessage({ id: 'page.upstream.step.healthyCheck.active.timeout' })}
      tooltip={formatMessage({ id: 'page.upstream.checks.active.timeout.description' })}
    >
      <Form.Item name={['checks', 'active', 'timeout']} noStyle initialValue={1}>
        <InputNumber disabled={readonly} min={0} />
      </Form.Item>
      <TimeUnit />
    </Form.Item>
  );
};

export default ActiveCheckTimeoutComponent;
