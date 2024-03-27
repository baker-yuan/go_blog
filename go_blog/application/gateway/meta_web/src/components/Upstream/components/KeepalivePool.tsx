import { Col, Form, InputNumber, Row } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  readonly?: boolean;
};

const KeepalivePool: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        label={formatMessage({ id: 'component.upstream.fields.keepalive_pool' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.keepalive_pool.tooltip' })}
      >
        <Row style={{ marginBottom: 10 }} gutter={10}>
          <Col span={5}>
            <Form.Item
              name={['keepalive_pool', 'size']}
              label={formatMessage({ id: 'component.upstream.fields.keepalive_pool.size' })}
              style={{ marginBottom: 0 }}
              initialValue={320}
            >
              <InputNumber
                min={1}
                placeholder={formatMessage({
                  id: 'component.upstream.fields.keepalive_pool.size.placeholder',
                })}
                disabled={readonly}
              />
            </Form.Item>
          </Col>
          <Col span={7}>
            <Form.Item
              name={['keepalive_pool', 'idle_timeout']}
              label={formatMessage({ id: 'component.upstream.fields.keepalive_pool.idle_timeout' })}
              style={{ marginBottom: 0 }}
              initialValue={60}
            >
              <InputNumber
                min={0}
                placeholder={formatMessage({
                  id: 'component.upstream.fields.keepalive_pool.idle_timeout.placeholder',
                })}
                disabled={readonly}
              />
            </Form.Item>
          </Col>
          <Col span={6}>
            <Form.Item
              name={['keepalive_pool', 'requests']}
              label={formatMessage({ id: 'component.upstream.fields.keepalive_pool.requests' })}
              style={{ marginBottom: 0 }}
              initialValue={1000}
            >
              <InputNumber
                min={1}
                placeholder={formatMessage({
                  id: 'component.upstream.fields.keepalive_pool.requests.placeholder',
                })}
                disabled={readonly}
              />
            </Form.Item>
          </Col>
        </Row>
      </Form.Item>
    </React.Fragment>
  );
};

export default KeepalivePool;
