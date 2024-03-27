import type { FormInstance } from 'antd/lib/form';
import React from 'react';

import PluginPage from '@/components/Plugin';

import Step1 from './Step1';

type Props = {
  form1: FormInstance;
  plugins: PluginComponent.Data;
};

const Page: React.FC<Props> = ({ form1, plugins }) => {
  return (
    <>
      <Step1 form={form1} disabled />
      <PluginPage initialData={plugins} readonly />
    </>
  );
};

export default Page;
