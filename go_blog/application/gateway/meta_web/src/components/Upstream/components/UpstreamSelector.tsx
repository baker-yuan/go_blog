import { Form, Select } from 'antd';
import React from 'react';
import { useIntl, useLocation } from 'umi';

type Upstream = {
  name?: string;
  id?: string;
};

type Props = {
  list?: Upstream[];
  disabled?: boolean;
  required?: boolean;
  onChange: (id: string) => void;
};

const UpstreamSelector: React.FC<Props> = ({ onChange, list = [], disabled, required }) => {
  const { formatMessage } = useIntl();
  const location = useLocation();

  return (
    <Form.Item
      label={formatMessage({ id: 'page.upstream.step.select.upstream' })}
      name="upstream_id"
    >
      <Select
        showSearch
        data-cy="upstream_selector"
        disabled={disabled}
        onChange={onChange}
        filterOption={(input, item) => item?.children.toLowerCase().includes(input.toLowerCase())}
      >
        <Select.Option value="None" disabled={required}>
          {formatMessage({ id: 'component.upstream.other.none' })}
        </Select.Option>
        {[
          {
            name: formatMessage({
              id: `page.upstream.step.select.upstream.select.option${
                !required && location.pathname === '/routes/create' ? '.serviceSelected' : ''
              }`,
            }),
            id: 'Custom',
          },
          ...list,
        ].map((item) => (
          <Select.Option value={item.id!} key={item.id}>
            {item.name}
          </Select.Option>
        ))}
      </Select>
    </Form.Item>
  );
};

export default UpstreamSelector;
