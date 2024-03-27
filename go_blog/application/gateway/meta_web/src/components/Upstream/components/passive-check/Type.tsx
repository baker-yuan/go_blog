import { Col, Form, Row, Select } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const options = [
  {
    label: 'HTTP',
    value: 'http',
  },
  {
    label: 'HTTPs',
    value: 'https',
  },
  {
    label: 'TCP',
    value: 'tcp',
  },
];

const PassiveCheckTypeComponent: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.Item
      label={formatMessage({ id: 'component.upstream.fields.checks.active.type' })}
      style={{ marginBottom: 0 }}
      tooltip={formatMessage({ id: 'component.upstream.fields.checks.active.type.tooltip' })}
    >
      <Row>
        <Col span={5}>
          <Form.Item name={['checks', 'passive', 'type']} initialValue="http">
            <Select disabled={readonly}>
              {options.map((item) => {
                return (
                  <Select.Option value={item.value} key={item.value}>
                    {item.label}
                  </Select.Option>
                );
              })}
            </Select>
          </Form.Item>
        </Col>
      </Row>
    </Form.Item>
  );
};

export default PassiveCheckTypeComponent;
