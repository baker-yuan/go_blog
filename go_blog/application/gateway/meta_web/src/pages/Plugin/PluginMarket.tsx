import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Card, notification } from 'antd';
import React, { useEffect, useState } from 'react';
import { useIntl } from 'umi';

import PluginPage from '@/components/Plugin';

import { createOrUpdate, fetchList } from './service';

const PluginMarket: React.FC = () => {
  const [initialData, setInitialData] = useState({});

  const initPageData = () => {
    fetchList().then(({ data }) => {
      const plugins: any = {};
      data.forEach(({ name, value }) => {
        plugins[name] = value;
      });
      setInitialData(plugins);
    });
  };

  useEffect(() => {
    initPageData();
  }, []);

  const { formatMessage } = useIntl();

  return (
    <PageHeaderWrapper title={formatMessage({ id: 'page.plugin.market.config' })}>
      <Card bordered={false}>
        <PluginPage
          initialData={initialData}
          type="global"
          schemaType="route"
          onChange={(pluginsData, pluginId, handleType) => {
            createOrUpdate({
              plugins: pluginsData,
            }).then(() => {
              initPageData();
              notification.success({
                message: formatMessage({
                  id: `page.plugin.${handleType}`,
                }),
              });
            });
          }}
        />
      </Card>
    </PageHeaderWrapper>
  );
};

export default PluginMarket;
