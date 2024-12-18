import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/', name: 'home', component: () => import('../views/Auth.vue'),
    },
    {
      path: '/manager', name: 'manager', component: () => import('../views/Manager.vue'), children: [
        { path: 'dashboard', name: 'dashboard', meta: { title: 'Dashboard' }, component: () => import('../views/Dashboard.vue') },
        {
          path: 'container', name: 'container', meta: { title: 'Container' }, component: () => import('../views/Container.vue'), children: [
            { path: 'creator', name: 'creator', meta: { title: 'Creator' }, component: () => import('../views/ContainerCreator.vue') }
          ]
        }
      ]
    },
    {
      path: '/auth', name: 'auth', meta: { title: 'Login' }, component: () => import('../views/Auth.vue'),
    }
  ],
})

export default router
