import { createRouter, createWebHistory } from 'vue-router';
import TikuComponent from '../components/TikuComponent/TikuComponent.vue';
import IntelligentImport from '../components/IntelligentImport/IntelligentImport.vue';

const routes = [
    { path: '/', component: TikuComponent },
    { path: '/import', component: IntelligentImport, name: 'import'},

];
  
  const router = createRouter({
    history: createWebHistory(),
    routes
  });
  
  export default router;