import { Settings as LayoutSettings } from '@ant-design/pro-layout';

const { REACT_APP_ENV, SERVE_ENV, SERVE_URL_DEV, SERVE_URL_TEST, CYPRESS_SERVE_ENV } = process.env;

const defaultSettings = {
  navTheme: 'dark',
  primaryColor: '#1890ff',
  layout: 'mix',
  contentWidth: 'Fluid',
  fixedHeader: false,
  autoHideHeader: false,
  fixSiderbar: true,
  colorWeak: false,
  menu: {
    locale: true,
  },
  title: 'APISIX Dashboard',
  pwa: false,
  iconfontUrl: '',
  serveUrlMap: {
    dev: SERVE_URL_DEV,
    test: SERVE_URL_TEST,
  },
  overwrite(env: Record<string, string>) {
    const { SERVE_URL_DEV, SERVE_URL_TEST } = env;
    this.serveUrlMap = {
      dev: SERVE_URL_DEV,
      test: SERVE_URL_TEST,
    };
    return this;
  },
} as LayoutSettings & {
  pwa: boolean;
  serveUrlMap: {
    dev: string;
    test: string;
  };
  overwrite: Function;
};

const { dev, test } = defaultSettings.serveUrlMap;
const throwPromptError = (message: TemplateStringsArray) => {
  throw new Error(
    `Please set '${message[0]}' in 'web/.env' file. Guide: https://apisix.apache.org/docs/dashboard/develop/#web`,
  );
};

console.log(defaultSettings.serveUrlMap);
const envs = [REACT_APP_ENV, SERVE_ENV, CYPRESS_SERVE_ENV];

if (envs.some((v) => v === 'test') && !test) {
  throwPromptError`SERVE_URL_TEST`;
}

if (envs.some((v) => v === 'dev') && !dev) {
  throwPromptError`SERVE_URL_DEV`;
}

export default defaultSettings;
