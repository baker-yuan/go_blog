import type { FormInstance } from 'antd';
import { AutoComplete, Form, Select } from 'antd';
import React, { useState } from 'react';
import { useIntl } from 'umi';

import { AlgorithmEnum, CommonHashKeyEnum, HashOnEnum } from '../constant';

type Props = {
  readonly?: boolean;
  form: FormInstance;
};

const CHash: React.FC<Props> = ({ form, readonly }) => {
  const { formatMessage } = useIntl();
  const [keySearchWord, setKeySearchWord] = useState('');

  const handleSearch = (search: string) => {
    setKeySearchWord(search);
  };
  return (
    <React.Fragment>
      <Form.Item
        name="hash_on"
        rules={[{ required: true }]}
        label={formatMessage({ id: 'component.upstream.fields.hash_on' })}
        tooltip={formatMessage({ id: 'component.upstream.fields.hash_on.tooltip' })}
        initialValue="vars"
      >
        <Select disabled={readonly}>
          {Object.entries(HashOnEnum).map(([label, value]) => (
            <Select.Option value={value} key={value}>
              {label}
            </Select.Option>
          ))}
        </Select>
      </Form.Item>
      {form.getFieldValue('hash_on') !== 'consumer' && (
        <Form.Item
          name="key"
          rules={[{ required: true }]}
          label={formatMessage({ id: 'component.upstream.fields.key' })}
          tooltip={formatMessage({ id: 'component.upstream.fields.key.tooltip' })}
          initialValue="remote_addr"
        >
          <AutoComplete disabled={readonly} onSearch={handleSearch}>
            {Object.entries(CommonHashKeyEnum)
              .filter(
                ([label, value]) =>
                  label.startsWith(keySearchWord) || value.startsWith(keySearchWord),
              )
              .map(([label, value]) => (
                <Select.Option value={value} key={value}>
                  {label}
                </Select.Option>
              ))}
          </AutoComplete>
        </Form.Item>
      )}
    </React.Fragment>
  );
};

const Component: React.FC<Props> = ({ readonly, form }) => {
  const { formatMessage } = useIntl();

  return (
    <React.Fragment>
      <Form.Item
        label={formatMessage({ id: 'page.upstream.step.type' })}
        name="type"
        rules={[{ required: true }]}
        initialValue="roundrobin"
      >
        <Select disabled={readonly}>
          {Object.entries(AlgorithmEnum).map(([label, value]) => {
            return (
              <Select.Option value={value} key={value}>
                {formatMessage({ id: `page.upstream.type.${label}` })}
              </Select.Option>
            );
          })}
        </Select>
      </Form.Item>
      <Form.Item shouldUpdate noStyle>
        {() => {
          if (form.getFieldValue('type') === 'chash') {
            return <CHash form={form} readonly={readonly} />;
          }
          return null;
        }}
      </Form.Item>
    </React.Fragment>
  );
};

export default Component;
