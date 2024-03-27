import { QuestionCircleOutlined } from '@ant-design/icons';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { Button, Card, Empty, Tooltip } from 'antd';
import React, { useEffect, useState } from 'react';
import { history, useIntl } from 'umi';

import { getGrafanaURL } from './service';

const Dashboard: React.FC = () => {
  const [grafanaURL, setGrafanaURL] = useState<string | undefined>();
  const { formatMessage } = useIntl();

  useEffect(() => {
    getGrafanaURL().then((url) => {
      setGrafanaURL(url);
    });
  }, []);

  return (
    <PageHeaderWrapper
      title={
        <>
          {formatMessage({ id: 'menu.dashboard' })}&nbsp;
          <Tooltip title={formatMessage({ id: 'page.dashboard.tip' })}>
            <QuestionCircleOutlined />
          </Tooltip>
        </>
      }
    >
      <Card>
        {!grafanaURL && (
          <Empty
            image="empty.svg"
            imageStyle={{
              height: 60,
            }}
            description={
              <span>
                {formatMessage({ id: 'page.dashboard.empty.description.grafanaNotConfig' })}
              </span>
            }
          >
            <Button
              type="primary"
              onClick={() => {
                history.replace({
                  pathname: '/settings',
                });
              }}
            >
              {formatMessage({ id: 'page.dashboard.button.grafanaConfig' })}
            </Button>
          </Empty>
        )}
        {grafanaURL && (
          <div>
            <iframe title="dashboard" src={grafanaURL} width="100%" height="860" frameBorder="0" />
          </div>
        )}
      </Card>
    </PageHeaderWrapper>
  );
};

export default Dashboard;
