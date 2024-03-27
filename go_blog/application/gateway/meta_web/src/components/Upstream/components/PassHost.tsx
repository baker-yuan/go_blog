import type { FormInstance } from 'antd';
import { Form, Input, notification, Select } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

type Props = {
  form: FormInstance;
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ form, readonly }) => {
  const { formatMessage } = useIntl();

  const options = [
    {
      value: 'pass',
      label: formatMessage({ id: 'page.upstream.step.pass-host.pass' }),
    },
    {
      value: 'node',
      label: formatMessage({ id: 'page.upstream.step.pass-host.node' }),
    },
    {
      value: 'rewrite',
      label: formatMessage({ id: 'page.upstream.step.pass-host.rewrite' }),
      disabled: true,
    },
  ];

  return (
    <React.Fragment>
      <Form.Item
        label={formatMessage({ id: 'page.upstream.step.pass-host' })}
        name="pass_host"
        initialValue="pass"
      >
        <Select disabled={readonly}>
          {options.map((item) => (
            <Select.Option value={item.value} key={item.value} disabled={item.disabled}>
              {item.label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item
        noStyle
        shouldUpdate={(prev, next) => {
          return prev.pass_host !== next.pass_host;
        }}
      >
        {() => {
          if (form.getFieldValue('pass_host') === 'rewrite') {
            return (
              <Form.Item
                label={formatMessage({ id: 'page.upstream.step.pass-host.upstream_host' })}
                name="upstream_host"
                rules={[
                  {
                    required: true,
                    message: '',
                  },
                ]}
              >
                <Input
                  disabled={readonly}
                  placeholder={formatMessage({ id: `page.upstream.upstream_host.required` })}
                />
              </Form.Item>
            );
          }

          if (
            form.getFieldValue('pass_host') === 'node' &&
            (form.getFieldValue('submitNodes') || []).length !== 1
          ) {
            notification.warning({
              message: formatMessage({
                id: 'component.upstream.other.pass_host-with-multiple-nodes.title',
              }),
              description: formatMessage({
                id: 'component.upstream.other.pass_host-with-multiple-nodes',
              }),
            });
            form.setFieldsValue({ pass_host: 'pass' });
          }
          return null;
        }}
      </Form.Item>
    </React.Fragment>
  );
};

export default Component;
