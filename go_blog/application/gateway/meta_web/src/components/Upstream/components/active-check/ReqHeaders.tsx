import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Col, Form, Input, Row } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import { removeBtnStyle } from '../../constant';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.List name={['checks', 'active', 'req_headers']}>
      {(fields, { add, remove }) => (
        <>
          <Form.Item
            label={formatMessage({ id: 'component.upstream.fields.checks.active.req_headers' })}
            tooltip={formatMessage({
              id: 'component.upstream.fields.checks.active.req_headers.tooltip',
            })}
            style={{ marginBottom: 0 }}
          >
            {fields.map((field, index) => (
              <Row style={{ marginBottom: 10 }} gutter={12} key={index}>
                <Col span={5}>
                  <Form.Item noStyle name={[field.name]}>
                    <Input
                      placeholder={formatMessage({
                        id: 'page.upstream.step.input.healthyCheck.active.req_headers',
                      })}
                      disabled={readonly}
                    />
                  </Form.Item>
                </Col>
                <Col style={{ ...removeBtnStyle, marginLeft: 0 }}>
                  {!readonly && fields.length > 0 && (
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
