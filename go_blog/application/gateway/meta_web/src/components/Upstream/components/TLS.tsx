import type { FormInstance } from 'antd';
import { Col, Form, Input, Row, Select } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  form: FormInstance;
  readonly?: boolean;
};

const TLSComponent: React.FC<Props> = ({ form, readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        label={formatMessage({ id: 'component.upstream.fields.tls' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.tls.tooltip' })}
        style={{ marginBottom: 0 }}
      >
        <Row>
          <Col span={5}>
            <Form.Item name={['custom', 'tls']} initialValue="disable">
              <Select disabled={readonly}>
                {['disable', 'enable'].map((item) => (
                  <Select.Option value={item} key={item}>
                    {formatMessage({ id: `component.global.${item}` })}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
          </Col>
        </Row>
      </Form.Item>
      <Form.Item
        noStyle
        shouldUpdate={(prev, next) => {
          return prev.custom.tls !== next.custom.tls;
        }}
      >
        {() => {
          if (form.getFieldValue(['custom', 'tls']) === 'enable') {
            return (
              <React.Fragment>
                <Form.Item
                  label={formatMessage({ id: 'component.upstream.fields.tls.client_cert' })}
                  name={['tls', 'client_cert']}
                  required
                  rules={[{ required: true, message: '' }, { max: 64 * 1024 }, { min: 128 }]}
                >
                  <Input.TextArea
                    disabled={readonly}
                    minLength={128}
                    maxLength={64 * 1024}
                    rows={5}
                    placeholder={formatMessage({
                      id: 'component.upstream.fields.tls.client_cert.required',
                    })}
                  />
                </Form.Item>
                <Form.Item
                  label={formatMessage({ id: 'component.upstream.fields.tls.client_key' })}
                  name={['tls', 'client_key']}
                  required
                  rules={[{ required: true, message: '' }, { max: 64 * 1024 }, { min: 128 }]}
                >
                  <Input.TextArea
                    disabled={readonly}
                    minLength={128}
                    maxLength={64 * 1024}
                    rows={5}
                    placeholder={formatMessage({
                      id: 'component.upstream.fields.tls.client_key.required',
                    })}
                  />
                </Form.Item>
              </React.Fragment>
            );
          }
          return null;
        }}
      </Form.Item>
    </React.Fragment>
  );
};

export default TLSComponent;
