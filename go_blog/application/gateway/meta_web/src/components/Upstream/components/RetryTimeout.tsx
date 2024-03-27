import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const RetryTimeout: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        name="retry_timeout"
        label={formatMessage({ id: 'component.upstream.fields.retry_timeout' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.retry_timeout.tooltip' })}
      >
        <InputNumber disabled={readonly} />
      </Form.Item>
    </React.Fragment>
  );
};

export default RetryTimeout;
