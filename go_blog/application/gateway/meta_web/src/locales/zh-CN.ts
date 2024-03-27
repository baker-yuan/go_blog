import { ActionBarZhCN } from '@/components/ActionBar';

import Plugin from '../components/Plugin/locales/zh-CN';
import PluginFlow from '../components/PluginFlow/locales/zh-CN';
import RawDataEditor from '../components/RawDataEditor/locales/zh-CN';
import UpstreamComponent from '../components/Upstream/locales/zh-CN';
import component from './zh-CN/component';
import globalHeader from './zh-CN/globalHeader';
import menu from './zh-CN/menu';
import other from './zh-CN/other';
import pwa from './zh-CN/pwa';
import settings from './zh-CN/setting';
import settingDrawer from './zh-CN/settingDrawer';

export default {
  'navBar.lang': '语言',
  'layout.user.link.help': '帮助',
  'layout.user.link.privacy': '隐私',
  'layout.user.link.terms': '条款',
  'app.preview.down.block': '下载此页面到本地项目',
  ...globalHeader,
  ...menu,
  ...settingDrawer,
  ...settings,
  ...pwa,
  ...component,
  ...other,
  ...ActionBarZhCN,
  ...Plugin,
  ...PluginFlow,
  ...RawDataEditor,
  ...UpstreamComponent,
};
