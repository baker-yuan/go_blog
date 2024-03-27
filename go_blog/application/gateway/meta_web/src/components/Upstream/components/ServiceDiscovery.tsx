import type { FormInstance } from 'antd';
import { Form, Input, Select } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import ServiceDiscoveryArgs from '@/components/Upstream/components/ServiceDiscoveryArgs';

type Props = {
  form: FormInstance;
  readonly?: boolean;
};

const discoveryType = {
  dns: {},
  consul_kv: {},
  nacos: {
    args: ['group_name', 'namespace_id'],
  },
  eureka: {},
  kubernetes: {},
};

const ServiceDiscovery: React.FC<Props> = ({ readonly, form }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        name="discovery_type"
        label={formatMessage({ id: 'component.upstream.fields.discovery_type' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.discovery_type.tooltip' })}
        rules={[{ required: true }]}
      >
        <Select
          disabled={readonly}
          placeholder={formatMessage({
            id: 'component.upstream.fields.discovery_type.placeholder',
          })}
        >
          {Object.keys(discoveryType).map((item) => {
            return (
              <Select.Option key={item} value={item}>
                {formatMessage({ id: `component.upstream.fields.discovery_type.type.${item}` })}
              </Select.Option>
            );
          })}
        </Select>
      </Form.Item>
      <Form.Item
        name="service_name"
        label={formatMessage({ id: 'component.upstream.fields.service_name' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.service_name.tooltip' })}
        rules={[{ required: true }, { min: 1 }, { max: 256 }]}
      >
        <Input
          disabled={readonly}
          placeholder={formatMessage({ id: 'component.upstream.fields.service_name.placeholder' })}
        />
      </Form.Item>
      <Form.Item shouldUpdate noStyle>
        {() => {
          if (!form.getFieldValue('discovery_type')) return null;

          const { args } = discoveryType[form.getFieldValue('discovery_type')];
          if (args && args.length > 0) {
            return <ServiceDiscoveryArgs readonly={readonly} args={args} />;
          }
          return null;
        }}
      </Form.Item>
    </React.Fragment>
  );
};

export default ServiceDiscovery;
