import { LockTwoTone, UserOutlined } from '@ant-design/icons';
import { Form, Input, Tooltip } from 'antd';
import type { FormInstance } from 'antd/lib/form';
import React from 'react';
import { formatMessage, request } from 'umi';

import type { UserModule } from '@/pages/User/typing';

const formRef = React.createRef<FormInstance>();

const LoginMethodPassword: UserModule.LoginMethod = {
  id: 'password',
  name: formatMessage({ id: 'component.user.loginMethodPassword' }),
  render: () => {
    return (
      <Form ref={formRef} name="control-ref">
        <Form.Item
          name="username"
          rules={[
            {
              required: true,
              message: formatMessage({ id: 'component.user.loginMethodPassword.inputUsername' }),
            },
          ]}
        >
          <Input
            size="large"
            type="text"
            placeholder={formatMessage({ id: 'component.user.loginMethodPassword.username' })}
            prefix={
              <UserOutlined
                style={{
                  color: '#1890ff',
                }}
              />
            }
          />
        </Form.Item>
        <Form.Item
          name="password"
          rules={[
            {
              required: true,
              message: formatMessage({ id: 'component.user.loginMethodPassword.inputPassword' }),
            },
          ]}
        >
          <Input
            size="large"
            type="password"
            placeholder={formatMessage({ id: 'component.user.loginMethodPassword.password' })}
            prefix={<LockTwoTone />}
          />
        </Form.Item>
        <Form.Item>
          <Tooltip
            title={formatMessage({ id: 'component.user.loginMethodPassword.modificationMethod' })}
          >
            <a
              href="https://github.com/apache/apisix-dashboard/blob/master/api/conf/conf.yaml#L70-L75"
              target="_blank"
            >
              {formatMessage({ id: 'component.user.loginMethodPassword.changeDefaultAccount' })}
            </a>
          </Tooltip>
        </Form.Item>
      </Form>
    );
  },
  getData(): UserModule.LoginData {
    if (formRef.current) {
      const data = formRef.current.getFieldsValue();
      return {
        username: data.username,
        password: data.password,
      };
    }
    return {};
  },
  checkData: async () => {
    if (formRef.current) {
      try {
        await formRef.current.validateFields();
        return true;
      } catch (e) {
        return false;
      }
    }
    return false;
  },
  submit: async ({ username, password }) => {
    if (username !== '' && password !== '') {
      try {
        const result = await request('/user/login', {
          method: 'POST',
          requestType: 'json',
          data: {
            username,
            password,
          },
        });

        localStorage.setItem('token', result.data.token);
        return {
          status: true,
          message: formatMessage({ id: 'component.user.loginMethodPassword.success' }),
          data: [],
        };
      } catch (e) {
        // NOTE: API failed, using errorHandler
        return {
          status: false,
          message: '',
          data: [],
        };
      }
    } else {
      return {
        status: false,
        message: formatMessage({ id: 'component.user.loginMethodPassword.fieldInvalid' }),
        data: [],
      };
    }
  },
  logout: () => {
    localStorage.removeItem('token');
  },
};

export default LoginMethodPassword;
