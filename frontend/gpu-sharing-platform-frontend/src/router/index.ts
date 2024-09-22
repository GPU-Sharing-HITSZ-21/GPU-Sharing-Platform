import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TerminalView from "@/views/TerminalView.vue";
import TestInstance from "@/views/TestInstance.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/terminal',
      name: 'terminal',
      component: TerminalView
    },
    {
      path: '/test_instance',
      name: 'test_instance',
      component: TestInstance
    }
  ]
})

export default router
