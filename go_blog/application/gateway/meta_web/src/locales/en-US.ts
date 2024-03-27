import { ActionBarEnUS } from '@/components/ActionBar';

import Plugin from '../components/Plugin/locales/en-US';
import PluginFlow from '../components/PluginFlow/locales/en-US';
import RawDataEditor from '../components/RawDataEditor/locales/en-US';
import UpstreamComponent from '../components/Upstream/locales/en-US';
import component from './en-US/component';
import globalHeader from './en-US/globalHeader';
import menu from './en-US/menu';
import other from './en-US/other';
import pwa from './en-US/pwa';
import settings from './en-US/setting';
import settingDrawer from './en-US/settingDrawer';

export default {
  'navBar.lang': 'Languages',
  'layout.user.link.help': 'Help',
  'layout.user.link.privacy': 'Privacy',
  'layout.user.link.terms': 'Terms',
  'app.preview.down.block': 'Download this page to your local project',
  ...globalHeader,
  ...menu,
  ...settingDrawer,
  ...settings,
  ...pwa,
  ...component,
  ...other,
  ...ActionBarEnUS,
  ...Plugin,
  ...PluginFlow,
  ...RawDataEditor,
  ...UpstreamComponent,
};
