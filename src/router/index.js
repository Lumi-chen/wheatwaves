import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import CustomLayout from '@/layout/CustomLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: '',
      component: CustomLayout,
      children: [
        {
          path: '/',
          name: '',
          component: HomeView,
        }
      ]
    },
  ]
})

export default router
