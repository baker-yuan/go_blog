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
      label={formatMessage({ id: 'component.upstream.fields.checks.active.unhealthy.timeouts' })}
      required
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.unhealthy.timeouts.tooltip',
      })}
    >
      <Form.Item name={['checks', 'active', 'unhealthy', 'timeouts']} noStyle initialValue={3}>
        <InputNumber disabled={readonly} min={1} max={254} />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
