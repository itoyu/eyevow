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
      name: 'aboutTutorial',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/AboutTutorial.vue')
    },
    {
      path: '/about/story',
      name: 'aboutStory',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStory.vue')
    },
    {
      path: '/about/story/:id', //:id(\\d+)
      name: 'aboutStoryDetail',
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStoryDetail.vue'),
      props: route => ({ 'id': Number(route.params.id) })
    },
    {
      path: '/about/statement',
      name: 'aboutStatement',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStatement.vue')
    },
    {
      path: '/about/credit',
      name: 'aboutCredit',
      component: () => import(/* webpackChunkName: "about" */ './views/AboutCredit.vue')
    },
    {
      path: '/vow',
      name: 'vow',
      component: () => import(/* webpackChunkName: "about" */ './views/Vow.vue')
    },
    {
      path: '/cheer',
      name: 'cheer',
      component: () => import(/* webpackChunkName: "about" */ './views/Cheer.vue')
    },
    {
      path: '/edit',
      name: 'edit',
      component: () => import(/* webpackChunkName: "about" */ './views/Edit.vue')
    },
    {
      path: '/product/:id(\\d+)',
      name: 'productDetail',
      component: () => import(/* webpackChunkName: "about" */ './views/Product.vue'),
      props: route => ({ 'id': Number(route.params.id) })
    }
  ]
})
