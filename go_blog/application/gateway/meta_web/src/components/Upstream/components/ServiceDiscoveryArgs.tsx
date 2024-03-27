import { Form, Input } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  args: string[];
  readonly?: boolean;
};

const ServiceDiscoveryArgs: React.FC<Props> = ({ readonly, args }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      {args.map((item) => {
        return (
          <Form.Item
            key={item}
            name={['discovery_args', item]}
            label={formatMessage({ id: `component.upstream.fields.discovery_args.${item}` })}
            tooltip={formatMessage({
              id: `component.upstream.fields.discovery_args.${item}.tooltip`,
            })}
          >
            <Input
              disabled={readonly}
              placeholder={formatMessage({
                id: `component.upstream.fields.discovery_args.${item}.placeholder`,
              })}
            />
          </Form.Item>
        );
      })}
    </React.Fragment>
  );
};

export default ServiceDiscoveryArgs;
