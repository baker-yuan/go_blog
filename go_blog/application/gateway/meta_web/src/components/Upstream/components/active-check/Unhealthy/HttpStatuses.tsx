import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Col, Form, InputNumber, Row } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import { removeBtnStyle } from '@/components/Upstream';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.List
      name={['checks', 'active', 'unhealthy', 'http_statuses']}
      initialValue={[429, 404, 500, 501, 502, 503, 504, 505]}
    >
      {(fields, { add, remove }) => (
        <>
          <Form.Item
            required
            label={formatMessage({ id: 'page.upstream.step.healthyCheck.passive.http_statuses' })}
            style={{ marginBottom: 0 }}
          >
            {fields.map((field, index) => (
              <Row style={{ marginBottom: 10 }} key={index}>
                <Col md={4} lg={4} xl={4} xxl={2}>
                  <Form.Item style={{ marginBottom: 0 }} name={[field.name]}>
                    <InputNumber disabled={readonly} min={200} max={599} />
                  </Form.Item>
                </Col>
                <Col style={removeBtnStyle}>
                  {!readonly && (
                    <MinusCircleOutlined
                      onClick={() => {
                        remove(field.name);
                      }}
                    />
                  )}
                </Col>
              </Row>
            ))}
          </Form.Item>
          {!readonly && (
            <Form.Item wrapperCol={{ offset: 3 }}>
              <Button type="dashed" onClick={() => add()}>
                <PlusOutlined />
                {formatMessage({
                  id: 'component.global.add',
                })}
              </Button>
            </Form.Item>
          )}
        </>
      )}
    </Form.List>
  );
};
export default Component;
