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
      label={formatMessage({ id: 'component.upstream.fields.retries' })}
      tooltip={formatMessage({ id: 'component.upstream.fields.retries.tooltip' })}
      name="retries"
    >
      <InputNumber disabled={readonly} />
    </Form.Item>
  );
};

export default Component;
