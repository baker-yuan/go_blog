import { MinusCircleOutlined } from '@ant-design/icons';
import { AutoComplete, Checkbox, Col, Form, Input, Row } from 'antd';
import React, { useState } from 'react';
import { useIntl } from 'umi';

import { HEADER_LIST } from '@/pages/Route/constants';

import styles from './index.less';

const { Option } = AutoComplete;

const DebugParamsView: React.FC<RouteModule.DebugViewProps> = (props) => {
  const { formatMessage } = useIntl();

  const allSelectOptions = props.inputType === 'header' ? HEADER_LIST : [];
  const [result, setResult] = useState<string[]>(allSelectOptions);

  const onSearch = (value: string) => {
    setResult(
      allSelectOptions.filter((option) => option.toLowerCase().startsWith(value.toLowerCase())),
    );
  };

  return (
    <Form name={props.name} className={styles.routeDebugDraw} form={props.form}>
      <Form.List name="params">
        {(fields, { add, remove }) => {
          return (
            <>
              {fields.map((field, index) => (
                <Row gutter={16} key={field.name}>
                  <Col span={1}>
                    <Form.Item
                      {...field}
                      name={[field.name, 'check']}
                      fieldKey={[field.fieldKey, 'check']}
                      style={{ textAlign: 'right' }}
                      valuePropName="checked"
                    >
                      {fields.length > 1 && index !== fields.length - 1 && <Checkbox />}
                    </Form.Item>
                  </Col>
                  <Col span={8}>
                    <Form.Item
                      {...field}
                      name={[field.name, 'key']}
                      fieldKey={[field.fieldKey, 'key']}
                    >
                      <AutoComplete
                        onSearch={onSearch}
                        placeholder={formatMessage({ id: 'page.route.input.placeholder.paramKey' })}
                        onChange={() => {
                          // only last line key field input can trigger add new line event
                          if (index === fields.length - 1) {
                            add();
                            const prevData = props.form.getFieldsValue();
                            // auto change the prev line checkbox checked
                            prevData.params[index].check = true;
                            props.form.setFieldsValue(prevData);
                          }
                        }}
                      >
                        {result.map((value) => (
                          <Option key={value} value={value}>
                            {value}
                          </Option>
                        ))}
                      </AutoComplete>
                    </Form.Item>
                  </Col>
                  <Col span={8}>
                    <Form.Item
                      {...field}
                      name={[field.name, 'value']}
                      fieldKey={[field.fieldKey, 'value']}
                    >
                      <Input
                        placeholder={formatMessage({
                          id: 'page.route.input.placeholder.paramValue',
                        })}
                      />
                    </Form.Item>
                  </Col>
                  <Col>
                    {fields.length > 1 && index !== fields.length - 1 && (
                      <MinusCircleOutlined onClick={() => remove(field.name)} />
                    )}
                  </Col>
                </Row>
              ))}
            </>
          );
        }}
      </Form.List>
    </Form>
  );
};

export default DebugParamsView;
