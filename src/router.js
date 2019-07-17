import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/views/Home.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/About.vue')
    },
    {
      path: '/product',
      name: 'product',
      component: () => import(/* webpackChunkName: "about" */ './views/ProductList.vue')
    },
    {
      path: '/product/:id(\\d+)',
      name: 'productDetail',
      component: () => import(/* webpackChunkName: "about" */ './views/Product.vue'),
      props: route => ({ 'id': Number(route.params.id) })
    }
  ]
})
