import './libs/iconfont';

import type { Settings as LayoutSettings } from '@ant-design/pro-layout';
import { isPlainObject } from 'lodash';
import React from 'react';
import type { RequestConfig } from 'umi';
import { history } from 'umi';

import Footer from '@/components/Footer';
import RightContent from '@/components/RightContent';
import { errorHandler, getMenuData, getUrlQuery } from '@/helpers';
import { queryCurrent } from '@/services/user';

import defaultSettings from '../config/defaultSettings';

export async function getInitialState(): Promise<{
  currentUser?: API.CurrentUser;
  settings?: LayoutSettings;
}> {
  const token = localStorage.getItem('token');
  if (!token) {
    const redirect = getUrlQuery('redirect') || '/';
    history.replace(`/user/login?redirect=${redirect}`);
  }

  const currentUser = await queryCurrent();
  return {
    currentUser,
    settings: defaultSettings,
  };
}

export const layout = ({ initialState }: { initialState: { settings?: LayoutSettings } }) => {
  return {
    headerRender: undefined,
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    footerRender: () => <Footer />,
    menuHeaderRender: undefined,
    menuDataRender: getMenuData,
    ...initialState?.settings,
  };
};

/* eslint no-param-reassign: ["error", { "props": true, "ignorePropertyModificationsFor": ["obj"] }] */
const nullValueFilter = (obj: Record<string, any>) => {
  Object.entries(obj).forEach(([key, value]) => {
    if (isPlainObject(value)) {
      nullValueFilter(value);
    } else if ([null, undefined].includes(value)) {
      delete obj[key];
    }
  });
};

export const request: RequestConfig = {
  prefix: '/apisix/admin',
  errorHandler,
  credentials: 'same-origin',
  requestInterceptors: [
    (url, options) => {
      const newOptions = { ...options };
      if (newOptions.data) {
        nullValueFilter(newOptions.data);
      }
      newOptions.headers = {
        ...options.headers,
        Authorization: localStorage.getItem('token') || '',
      };
      return {
        url,
        options: { ...newOptions, interceptors: true },
      };
    },
  ],
  responseInterceptors: [
    async (res) => {
      if (!res.ok) {
        // NOTE: http code >= 400, using errorHandler
        return res;
      }

      const data = await res.json();
      const { code = -1 } = data as Res<any>;
      if (code !== 0) {
        // eslint-disable-next-line
        return Promise.reject({ response: res, data });
      }
      return data;
    },
  ],
};
