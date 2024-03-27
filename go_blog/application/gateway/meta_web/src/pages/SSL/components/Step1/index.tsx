import { Form, Select } from 'antd';
import type { FormInstance } from 'antd/es/form';
import type { UploadFile } from 'antd/lib/upload/interface';
import React, { useState } from 'react';
import { useIntl } from 'umi';

import CertificateForm from '@/pages/SSL/components/CertificateForm';
import type { UploadType } from '@/pages/SSL/components/CertificateUploader';
import CertificateUploader from '@/pages/SSL/components/CertificateUploader';

type CreateType = 'Upload' | 'Input';

type Props = {
  form: FormInstance;
};

const Step: React.FC<Props> = ({ form }) => {
  const [publicKeyList, setPublicKeyList] = useState<UploadFile[]>([]);
  const [privateKeyList, setPrivateKeyList] = useState<UploadFile[]>([]);

  const [createType, setCreateType] = useState<CreateType>('Input');

  const { formatMessage } = useIntl();

  const onRemove = (type: UploadType) => {
    if (type === 'PUBLIC_KEY') {
      form.setFieldsValue({
        cert: '',
        sni: '',
        expireTime: undefined,
      });
      setPublicKeyList([]);
    } else {
      form.setFieldsValue({ key: '' });
      setPrivateKeyList([]);
    }
  };

  const handleSuccess = ({
    cert,
    key,
    ...rest
  }: Partial<SSLModule.UploadPrivateSuccessData & SSLModule.UploadPublicSuccessData>) => {
    if (cert) {
      setPublicKeyList(rest.publicKeyList!);
      form.setFieldsValue({ cert });
    } else {
      form.setFieldsValue({ key });
      setPrivateKeyList(rest.privateKeyList!);
    }
  };
  return (
    <>
      <Form.Item label={formatMessage({ id: 'page.ssl.form.itemLabel.way' })} required>
        <Select
          placeholder={formatMessage({ id: 'page.ssl.select.placeholder.selectCreateWays' })}
          defaultValue="Input"
          onChange={(value: CreateType) => {
            form.setFieldsValue({
              key: '',
              cert: '',
              sni: '',
              expireTime: undefined,
            });
            setCreateType(value);
          }}
          style={{ width: 100 }}
        >
          <Select.Option value="Input">
            {formatMessage({ id: 'page.ssl.selectOption.input' })}
          </Select.Option>
          <Select.Option value="Upload">{formatMessage({ id: 'page.ssl.upload' })}</Select.Option>
        </Select>
      </Form.Item>
      <div style={createType === 'Input' ? {} : { display: 'none' }}>
        <CertificateForm mode="EDIT" form={form} />
      </div>
      {Boolean(createType === 'Upload') && (
        <CertificateUploader
          onSuccess={handleSuccess}
          onRemove={onRemove}
          data={{ publicKeyList, privateKeyList }}
        />
      )}
    </>
  );
};
export default Step;
