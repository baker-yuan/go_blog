import { Form, Input } from 'antd';
import type { FormInstance } from 'antd/lib/form';
import React from 'react';
import { useIntl } from 'umi';

import UpstreamForm from '@/components/Upstream';

type Props = {
  form: FormInstance;
  disabled?: boolean;
  upstreamRef?: React.MutableRefObject<any>;
  neverReadonly?: boolean;
};

const Step1: React.FC<Props> = ({ form, disabled, upstreamRef, neverReadonly }) => {
  const { formatMessage } = useIntl();

  return (
    <>
      <Form labelCol={{ span: 3 }} form={form}>
        <Form.Item
          label={formatMessage({ id: 'page.upstream.step.name' })}
          name="name"
          rules={[
            {
              required: true,
              message: formatMessage({ id: 'page.upstream.step.input.upstream.name' }),
            },
          ]}
        >
          <Input
            placeholder={formatMessage({ id: 'page.upstream.step.input.upstream.name' })}
            disabled={disabled}
          />
        </Form.Item>
        <Form.Item label={formatMessage({ id: 'page.upstream.step.description' })} name="desc">
          <Input.TextArea
            placeholder={formatMessage({ id: 'page.upstream.step.input.description' })}
            disabled={disabled}
          />
        </Form.Item>
      </Form>
      <UpstreamForm
        ref={upstreamRef}
        form={form}
        disabled={disabled}
        neverReadonly={neverReadonly}
      />
    </>
  );
};

export default Step1;
