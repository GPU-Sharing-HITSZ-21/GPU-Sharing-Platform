import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TerminalView from "@/views/TerminalView.vue";
import TestInstance from "@/views/TestInstance.vue";
import LoginRegister from "@/views/LoginRegister.vue";
import UserPodsView from "@/views/UserPodsView.vue";
import CreatePodView from "@/views/CreatePodView.vue";
import FileView from "@/views/FileView.vue";
import jobView from "@/views/JobView.vue";
import JobView from "@/views/JobView.vue";
import HomePageView from "@/views/HomePageView.vue";

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
    },
    {
      path: '/auth',
      name: 'auth',
      component: LoginRegister
    },
    {
      path: '/pods',
      name: 'pods',
      component: UserPodsView
    },
    {
      path: '/create-pod',
      name: 'create-pod',
      component: CreatePodView
    },
    {
      path: '/fileupload',
      name: 'fileupload',
      component: FileView
    },
    {
      path: '/job',
      name: 'job',
      component: JobView
    },
    {
      path: '/homepage',
      name: 'homepage',
      component: HomePageView
    }
  ]
})

export default router
