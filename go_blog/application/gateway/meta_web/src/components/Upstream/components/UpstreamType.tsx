import { Form, Select } from 'antd';
import type { FormInstance } from 'antd/lib/form';
import React from 'react';
import { useIntl } from 'umi';

import Nodes from '@/components/Upstream/components/Nodes';
import ServiceDiscovery from '@/components/Upstream/components/ServiceDiscovery';

type Props = {
  form: FormInstance;
  readonly?: boolean;
};

const UpstreamType: React.FC<Props> = ({ readonly, form }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        label={formatMessage({ id: 'component.upstream.fields.upstream_type' })}
        name="upstream_type"
        rules={[{ required: true }]}
        initialValue="node"
      >
        <Select disabled={readonly}>
          <Select.Option value="node">
            {formatMessage({ id: 'component.upstream.fields.upstream_type.node' })}
          </Select.Option>
          <Select.Option value="service_discovery">
            {formatMessage({ id: 'component.upstream.fields.upstream_type.service_discovery' })}
          </Select.Option>
        </Select>
      </Form.Item>

      <Form.Item shouldUpdate noStyle>
        {() => {
          if (form.getFieldValue('upstream_type') === 'node') {
            return <Nodes readonly={readonly} />;
          }
          return <ServiceDiscovery form={form} readonly={readonly} />;
        }}
      </Form.Item>
    </React.Fragment>
  );
};

export default UpstreamType;
