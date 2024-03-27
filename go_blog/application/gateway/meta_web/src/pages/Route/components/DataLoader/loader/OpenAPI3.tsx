import { Col, Form, Row, Switch } from 'antd';
import React, { memo } from 'react';

import { useIntl } from '@@/plugin-locale/localeExports';

const DataLoaderOpenAPI3: React.FC = () => {
  const { formatMessage } = useIntl();

  return (
    <Row gutter={16}>
      <Col span={12}>
        <Form.Item
          name="merge_method"
          label={formatMessage({ id: 'page.route.data_loader.labels.openapi3_merge_method' })}
          tooltip={formatMessage({ id: 'page.route.data_loader.tips.openapi3_merge_method' })}
          initialValue={true}
        >
          <Switch defaultChecked />
        </Form.Item>
      </Col>
    </Row>
  );
};

export default memo(DataLoaderOpenAPI3);
