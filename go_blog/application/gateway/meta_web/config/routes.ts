const routes = [
  {
    path: '/',
    component: './Dashboard',
  },
  {
    path: '/dashboard',
    component: './Dashboard',
  },
  {
    path: '/serverinfo',
    component: './ServerInfo',
  },
  {
    path: '/routes/list',
    component: './Route/List',
  },
  {
    path: '/routes/create',
    component: './Route/Create',
  },
  {
    path: '/routes/:rid/edit',
    component: './Route/Create',
  },
  {
    path: '/routes/:rid/duplicate',
    component: './Route/Create',
  },
  {
    path: '/ssl/:id/edit',
    component: './SSL/Create',
  },
  {
    path: '/ssl/list',
    component: './SSL/List',
  },
  {
    path: '/ssl/create',
    component: './SSL/Create',
  },
  {
    path: '/upstream/list',
    component: './Upstream/List',
  },
  {
    path: '/upstream/create',
    component: './Upstream/Create',
  },
  {
    path: '/upstream/:id/edit',
    component: './Upstream/Create',
  },
  {
    path: '/consumer/list',
    component: './Consumer/List',
  },
  {
    path: '/consumer/create',
    component: './Consumer/Create',
  },
  {
    path: '/consumer/:username/edit',
    component: './Consumer/Create',
  },
  {
    path: '/plugin/list',
    component: './Plugin/List',
  },
  {
    path: '/plugin/market',
    component: './Plugin/PluginMarket',
  },
  {
    path: '/service/list',
    component: './Service/List',
  },
  {
    path: '/service/create',
    component: './Service/Create',
  },
  {
    path: '/service/:serviceId/edit',
    component: './Service/Create',
  },
  {
    path: '/proto/list',
    component: './Proto/List',
  },
  {
    path: '/settings',
    component: './Setting',
  },
  {
    path: '/plugin-template/list',
    component: './PluginTemplate/List',
  },
  {
    path: 'plugin-template/create',
    component: './PluginTemplate/Create',
  },
  {
    path: '/plugin-template/:id/edit',
    component: './PluginTemplate/Create',
  },
  {
    path: '/user/login',
    component: './User/Login',
    layout: false,
  },
  {
    path: '/user/logout',
    component: './User/Logout',
    layout: false,
  },
  {
    component: './404',
  },
];

export default routes;
