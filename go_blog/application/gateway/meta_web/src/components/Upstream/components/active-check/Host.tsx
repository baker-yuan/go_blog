import { Form, Input } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();
  return (
    <Form.Item
      label={formatMessage({ id: 'component.upstream.fields.checks.active.host' })}
      tooltip={formatMessage({ id: 'component.upstream.fields.checks.active.host.tooltip' })}
      style={{ marginBottom: 0 }}
    >
      <Form.Item
        name={['checks', 'active', 'host']}
        rules={[
          {
            pattern: new RegExp(/^\*?[0-9a-zA-Z-._]+$/, 'g'),
            message: formatMessage({ id: 'component.upstream.fields.checks.active.host.scope' }),
          },
        ]}
      >
        <Input
          placeholder={formatMessage({
            id: 'component.upstream.fields.checks.active.host.required',
          })}
          disabled={readonly}
        />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
