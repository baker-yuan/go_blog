import { defineConfig } from 'umi';

import defaultSettings from './defaultSettings';
import proxy from './proxy';
import routes from './routes';

const { REACT_APP_ENV, SERVE_URL_DEV, SERVE_URL_TEST } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  locale: {
    default: 'en-US',
    antd: true,
    baseNavigator: true,
  },
  dynamicImport: {
    loading: '@/components/PageLoading/index',
  },
  targets: {
    ie: 11,
  },
  routes,
  layout: {
    name: 'APISIX Dashboard',
    locale: true,
    logo: '/favicon.png',
  },
  base: '/',
  publicPath: '/',
  define: {
    REACT_APP_ENV: REACT_APP_ENV || false,
    'process.env': {
      SERVE_URL_DEV,
      SERVE_URL_TEST,
    },
  },
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    'primary-color': defaultSettings.primaryColor,
  },
  // @ts-ignore
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
  outputPath: '../output/html',
  extraBabelPlugins: [
    [
      'babel-plugin-istanbul',
      {
        exclude: ['**/.umi', '**/locales'],
      },
    ],
  ],
});
