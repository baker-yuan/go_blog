import defaultSettings from './defaultSettings';

const { SERVE_ENV = 'dev' } = process.env;

export default {
  dev: {
    '/apisix/admin': {
      // NOTE: This is the manager-api pre-deployed in Azure just for preview, please refer to https://www.yuque.com/umijs/umi/proxy for more info.
      target: defaultSettings.serveUrlMap[SERVE_ENV],
      changeOrigin: true,
    },
  },
};
