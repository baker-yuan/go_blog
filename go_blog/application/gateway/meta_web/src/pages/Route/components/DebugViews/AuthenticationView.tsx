import { Card, Form, Input, Radio } from 'antd';
import React, { useState } from 'react';
import { useIntl } from 'umi';

import { AUTH_LIST } from '../../constants';
import styles from './index.less';

const AuthenticationView: React.FC<RouteModule.DebugViewProps> = (props) => {
  const { formatMessage } = useIntl();
  const [authType, setAuthType] = useState('none');

  const getAuthFormItems = () => {
    switch (props.form.getFieldValue('authType')) {
      case 'basic-auth':
        return (
          <>
            <Form.Item
              label={formatMessage({ id: 'page.route.form.itemLabel.username' })}
              name="username"
              rules={[
                {
                  required: true,
                  message: `${formatMessage({
                    id: 'component.global.pleaseEnter',
                  })}${formatMessage({ id: 'page.route.form.itemLabel.username' })}`,
                },
              ]}
            >
              <Input />
            </Form.Item>
            <Form.Item
              label={formatMessage({ id: 'page.route.form.itemLabel.password' })}
              name="password"
              rules={[
                {
                  required: true,
                  message: `${formatMessage({
                    id: 'component.global.pleaseEnter',
                  })}${formatMessage({ id: 'page.route.form.itemLabel.password' })}`,
                },
              ]}
            >
              <Input.Password />
            </Form.Item>
          </>
        );
      case 'jwt-auth':
        return (
          <Form.Item
            label={formatMessage({ id: 'page.route.form.itemLabel.token' })}
            name="Authorization"
            rules={[
              {
                required: true,
                message: `${formatMessage({ id: 'component.global.pleaseEnter' })}${formatMessage({
                  id: 'page.route.form.itemLabel.token',
                })}`,
              },
            ]}
          >
            <Input />
          </Form.Item>
        );
      case 'key-auth':
        return (
          <Form.Item
            label={formatMessage({ id: 'page.route.form.itemLabel.apikey' })}
            name="apikey"
            rules={[
              {
                required: true,
                message: `${formatMessage({ id: 'component.global.pleaseEnter' })}${formatMessage({
                  id: 'page.route.form.itemLabel.apikey',
                })}`,
              },
            ]}
          >
            <Input />
          </Form.Item>
        );
      default:
        return <div>{formatMessage({ id: 'page.route.debugWithoutAuth' })}</div>;
    }
  };

  return (
    <Card>
      <Form name="authForm" form={props.form}>
        <div className={styles.authForm}>
          <Form.Item name="authType">
            <Radio.Group
              defaultValue={authType}
              onChange={(event) => {
                const currentValue = event.target.value;
                setAuthType(currentValue);
                props.form.setFieldsValue({ autyType: currentValue });
              }}
            >
              <Radio value="none">None</Radio>
              {AUTH_LIST.map((type) => (
                <Radio data-cy={type} key={type} value={type}>
                  {type}
                </Radio>
              ))}
            </Radio.Group>
          </Form.Item>
          <div>{getAuthFormItems()}</div>
        </div>
      </Form>
    </Card>
  );
};

export default AuthenticationView;
