import { Empty } from 'antd';
import type { FormInstance } from 'antd/es/form';
import React from 'react';
import { useIntl } from 'umi';

import ApiBreaker from './api-breaker';
import BasicAuth from './basic-auth';
import Cors from './cors';
import LimitConn from './limit-conn';
import LimitCount from './limit-count';
import LimitReq from './limit-req';
import ProxyMirror from './proxy-mirror';
import RefererRestriction from './referer-restriction';

type Props = {
  name: string;
  schema: Record<string, any> | undefined;
  form: FormInstance;
  renderForm: boolean;
};

export const PLUGIN_UI_LIST = [
  'api-breaker',
  'basic-auth',
  'cors',
  'limit-req',
  'limit-conn',
  'proxy-mirror',
  'referer-restriction',
  'limit-count',
];

export const PluginForm: React.FC<Props> = ({ name, schema, renderForm, form }) => {
  const { formatMessage } = useIntl();

  if (!renderForm) {
    return (
      <Empty
        style={{ marginTop: 100 }}
        description={formatMessage({ id: 'component.plugin.noConfigurationRequired' })}
      />
    );
  }

  switch (name) {
    case 'api-breaker':
      return <ApiBreaker form={form} schema={schema} />;
    case 'basic-auth':
      return <BasicAuth form={form} schema={schema} />;
    case 'limit-count':
      return <LimitCount form={form} schema={schema} />;
    case 'cors':
      return <Cors form={form} schema={schema} />;
    case 'limit-req':
      return <LimitReq form={form} schema={schema} />;
    case 'proxy-mirror':
      return <ProxyMirror form={form} schema={schema} />;
    case 'limit-conn':
      return <LimitConn form={form} schema={schema} />;
    case 'referer-restriction':
      return <RefererRestriction form={form} schema={schema} />;
    default:
      return null;
  }
};
