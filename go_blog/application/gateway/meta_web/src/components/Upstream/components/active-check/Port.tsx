import { Form, InputNumber } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();
  return (
    <Form.Item label={formatMessage({ id: 'component.upstream.fields.checks.active.port' })}>
      <Form.Item name={['checks', 'active', 'port']} noStyle>
        <InputNumber
          placeholder={formatMessage({
            id: 'component.upstream.fields.checks.active.port',
          })}
          disabled={readonly}
          min={1}
          max={65535}
        />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
