import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';
import React from 'react';

export default () => (
  <DefaultFooter
    copyright={`${new Date().getFullYear()} Apache APISIX`}
    links={[
      {
        key: 'GitHub',
        title: <GithubOutlined />,
        href: 'https://github.com/apache/apisix',
        blankTarget: true,
      },
    ]}
  />
);
