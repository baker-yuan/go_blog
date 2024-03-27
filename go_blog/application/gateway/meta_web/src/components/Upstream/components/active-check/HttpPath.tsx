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
      label={formatMessage({ id: 'component.upstream.fields.checks.active.http_path' })}
      tooltip={formatMessage({
        id: 'component.upstream.fields.checks.active.http_path.tooltip',
      })}
    >
      <Form.Item name={['checks', 'active', 'http_path']} noStyle initialValue="/">
        <Input
          disabled={readonly}
          placeholder={formatMessage({
            id: 'component.upstream.fields.checks.active.http_path.placeholder',
          })}
        />
      </Form.Item>
    </Form.Item>
  );
};

export default Component;
