import { Form } from 'antd';
import React from 'react';

import { FORM_ITEM_LAYOUT } from '@/pages/Route/constants';

import styles from '../../Create.less';
import MatchingRulesView from './MatchingRulesView';
import MetaView from './MetaView';
import ProxyRewrite from './ProxyRewrite';
import RequestConfigView from './RequestConfigView';

const Step1: React.FC<RouteModule.Step1PassProps> = (props) => {
  return (
    <>
      <Form {...FORM_ITEM_LAYOUT} form={props.form} layout="horizontal" className={styles.stepForm}>
        <MetaView {...props} />
        <RequestConfigView {...props} />
        <ProxyRewrite {...props} />
      </Form>
      <MatchingRulesView {...props} />
    </>
  );
};

export default Step1;
