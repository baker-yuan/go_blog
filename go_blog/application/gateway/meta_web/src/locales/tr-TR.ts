import { ActionBarEnUS } from '@/components/ActionBar';

import Plugin from '../components/Plugin/locales/tr-TR';
import PluginFlow from '../components/PluginFlow/locales/tr-TR';
import RawDataEditor from '../components/RawDataEditor/locales/tr-TR';
import UpstreamComponent from '../components/Upstream/locales/tr-TR';
import component from './tr-TR/component';
import globalHeader from './tr-TR/globalHeader';
import menu from './tr-TR/menu';
import other from './tr-TR/other';
import pwa from './tr-TR/pwa';
import settings from './tr-TR/setting';
import settingDrawer from './tr-TR/settingDrawer';

export default {
  'navBar.lang': 'Dil Seçenekleri',
  'layout.user.link.help': 'Yardım',
  'layout.user.link.privacy': 'Gizlilik',
  'layout.user.link.terms': 'Kurallar',
  'app.preview.down.block': 'Sayfayı yerel projenize indirin',
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
