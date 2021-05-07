export default [
  {
    path: '/signin',
    layout: false,
    name: 'signin',
    hideInMenu: true,
    component: './account/signin',
  },
  {
    path: '/welcome',
    name: 'welcome',
    icon: 'smile',
    component: './Welcome',
  },
  {
    path: '/table',
    name: 'list',
    icon: 'smile',
    component: './table',
  },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    component: './exception/404',
  },
];
