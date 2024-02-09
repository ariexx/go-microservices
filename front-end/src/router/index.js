import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/product/:id',
      name: 'productDetail',
      component: () => import('../views/product/DetailView.vue'),
      props: true
    },
    {
      path: '/payment/:orderId',
      name: 'payment',
      component: () => import('../views/payment/PaymentView.vue'),
      props: true
    }
  ]
})

export default router
