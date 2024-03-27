import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const ConcurrencyComponent: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.Item
      label={formatMessage({ id: 'component.upstream.fields.checks.active.concurrency' })}
      tooltip={formatMessage({ id: 'component.upstream.fields.checks.active.concurrency.tooltip' })}
    >
      <Form.Item name={['checks', 'active', 'concurrency']} noStyle initialValue={10}>
        <InputNumber disabled={readonly} min={0} />
      </Form.Item>
    </Form.Item>
  );
};

export default ConcurrencyComponent;
