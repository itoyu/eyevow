import Vue from 'vue'
import Router from 'vue-router'
import store from './store';
import Home from '@/views/Home.vue'

Vue.use(Router)

// export default new Router({
const router  = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta: { title:'Home' }
    },
    {
      path: '/about',
      name: 'aboutTutorial',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('./views/AboutTutorial.vue'),
      meta: { title:'About' }
    },
    {
      path: '/about/story',
      name: 'aboutStory',
      component: () => import('./views/AboutStory.vue'),
      meta: { title:'Story' }
    },
    {
      path: '/about/story/:id', //:id(\\d+)
      name: 'aboutStoryDetail',
      component: () => import('./views/AboutStoryDetail.vue'),
      props: route => ({ 'id': Number(route.params.id) }),
      meta: route => ({ title: `Story ${route.params.id}` })
    },
    {
      path: '/about/statement',
      name: 'aboutStatement',
      component: () => import('./views/AboutStatement.vue'),
      meta: { title:'Statement' }
    },
    {
      path: '/about/credit',
      name: 'aboutCredit',
      component: () => import('./views/AboutCredit.vue'),
      meta: { title:'Credit' }
    },
    {
      path: '/vow',
      name: 'vow',
      component: () => import('./views/Vow.vue'),
      meta: { title:'Vow' }
    },
    {
      path: '/achieve',
      name: 'achieve',
      component: () => import('./views/Achieve.vue'),
      meta: { title:'Achieved' }
    },
    {
      path: '/achieved',
      name: 'achieved',
      component: () => import('./views/Achieved.vue'),
      meta: { title:'Achieved' }
    },
    {
      path: '/eyevow',
      name: 'eyevow',
      component: () => import('./views/Eyevow.vue'),
      meta: { title:'eyevow' }
    },
    {
      path: '/cheer',
      name: 'cheer',
      component: () => import('./views/Cheer.vue'),
      meta: { title:'Cheer' }
    },
    {
      path: '/edit',
      name: 'edit',
      component: () => import('./views/Edit.vue'),
      meta: { title:'Edit' }
    },
    {
      path: '/product/:id(\\d+)',
      name: 'productDetail',
      component: () => import('./views/Product.vue'),
      props: route => ({ 'id': Number(route.params.id) }),
      meta: route => ({ title: `Product ${route.params.id}` })
    }
  ],
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})
export default router;

// router.afterEach((to,from) =>{
router.afterEach((to) => {
  // console.log('呼ばれたよ!')
  if (to.meta && to.meta.title) {
    store.dispatch('changePage', to.meta.title)
  }
})
