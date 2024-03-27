import type { FormInstance } from 'antd/lib/form';
import React from 'react';
import { useIntl } from 'umi';

import PluginPage from '@/components/Plugin';
import PluginFlow from '@/components/PluginFlow';

import Step1 from '../Step1';
import Step2 from '../Step2';

type Props = {
  form1: FormInstance;
  form2: FormInstance;
  redirect?: boolean;
  step3Data: RouteModule.Step3Data;
  advancedMatchingRules: RouteModule.MatchingRule[];
  upstreamRef: any;
  isEdit?: boolean;
};

const style = {
  marginTop: '40px',
};

const CreateStep4: React.FC<Props> = ({ form1, form2, redirect, upstreamRef, ...rest }) => {
  const { formatMessage } = useIntl();
  const { plugins = {}, plugin_config_id = '', script = {} } = rest.step3Data;

  return (
    <>
      <h2>{formatMessage({ id: 'page.route.steps.stepTitle.defineApiRequest' })}</h2>
      <Step1 {...rest} form={form1} disabled isEdit />
      {!redirect && (
        <>
          <h2 style={style}>
            {formatMessage({ id: 'page.route.steps.stepTitle.defineApiBackendServe' })}
          </h2>
          <Step2
            form={form2}
            upstreamRef={upstreamRef}
            disabled
            hasServiceId={form1.getFieldValue('service_id') !== ''}
          />
          <h2 style={style}>
            {formatMessage({ id: 'component.global.steps.stepTitle.pluginConfig' })}
          </h2>
          {Boolean(Object.keys(plugins).length !== 0 || plugin_config_id !== '') && (
            <PluginPage
              referPage="route"
              initialData={plugins}
              plugin_config_id={plugin_config_id}
              showSelector
              readonly
            />
          )}
          {Boolean(Object.keys(script || {}).length !== 0) && (
            <PluginFlow chart={script.chart} readonly />
          )}
        </>
      )}
    </>
  );
};

export default CreateStep4;
