import { createRouter, createWebHistory } from 'vue-router';

import AdapterLoginLayout from '../layout/LoginLayout.vue';
import AdapterLoginPage from '../pages/Login/index.vue'

import TikuAdapterLayout from '../layout/TikuAdapterLayout.vue';
import TikuComponent from '../pages/TikuComponent/index.vue';
import IntelligentImport from '../pages/IntelligentImport/index.vue';
import UserList from '../pages/UserList/index.vue'
import LogList from '../pages/LogList/index.vue'
const routes = [
  {
    path: '/',
    name: 'loginlayout',
    component: AdapterLoginLayout,
    redirect: 'login',
    children: [
      {
        path: 'login',
        name: 'adapterlogin',
        component: AdapterLoginPage
      }
    ]
  },
  {
    path: '/adapter',
    name: 'adapterlayout',
    component: TikuAdapterLayout,
    children: [
      {
        path: 'component',
        name: 'component',
        component: TikuComponent,
      },
      {
        path: 'import',
        name: 'import',
        component: IntelligentImport,
      },
      {
        path: 'userlist',
        name: 'userlist',
        component: UserList,
      },
      {
        path: 'loglist',
        name: 'loglist',
        component: LogList,
      }
    ]
  }

];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;


