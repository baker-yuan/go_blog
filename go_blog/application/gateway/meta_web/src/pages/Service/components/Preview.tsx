import type { FormInstance } from 'antd/lib/form';
import React from 'react';

import PluginPage from '@/components/Plugin';

import Step1 from './Step1';

type Props = {
  form: FormInstance;
  upstreamForm: FormInstance;
  plugins: PluginComponent.Data;
  upstreamRef: React.MutableRefObject<any>;
};

const Page: React.FC<Props> = ({ form, plugins, upstreamForm, upstreamRef }) => {
  return (
    <>
      <Step1 form={form} upstreamForm={upstreamForm} upstreamRef={upstreamRef} disabled />
      <PluginPage initialData={plugins} readonly />
    </>
  );
};

export default Page;
