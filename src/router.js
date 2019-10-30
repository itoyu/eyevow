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
      component: () => import(/* webpackChunkName: "about" */ './views/AboutTutorial.vue'),
      meta: { title:'About' }
    },
    {
      path: '/about/story',
      name: 'aboutStory',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStory.vue'),
      meta: { title:'Story' }
    },
    {
      path: '/about/story/:id', //:id(\\d+)
      name: 'aboutStoryDetail',
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStoryDetail.vue'),
      props: route => ({ 'id': Number(route.params.id) }),
      meta: route => ({ title: `Story ${route.params.id}` })
    },
    {
      path: '/about/statement',
      name: 'aboutStatement',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import(/* webpackChunkName: "about" */ './views/AboutStatement.vue'),
      meta: { title:'Statement' }
    },
    {
      path: '/about/credit',
      name: 'aboutCredit',
      component: () => import(/* webpackChunkName: "about" */ './views/AboutCredit.vue'),
      meta: { title:'Credit' }
    },
    {
      path: '/vow',
      name: 'vow',
      component: () => import(/* webpackChunkName: "about" */ './views/Vow.vue'),
      meta: { title:'Vow' }
    },
    {
      path: '/eyevow',
      name: 'eyevow',
      component: () => import(/* webpackChunkName: "about" */ './views/Eyevow.vue'),
      meta: { title:'eyevow' }
    },
    {
      path: '/cheer',
      name: 'cheer',
      component: () => import(/* webpackChunkName: "about" */ './views/Cheer.vue'),
      meta: { title:'Cheer' }
    },
    {
      path: '/edit',
      name: 'edit',
      component: () => import(/* webpackChunkName: "about" */ './views/Edit.vue'),
      meta: { title:'Edit' }
    },
    {
      path: '/edit/review',
      name: 'editReview',
      component: () => import(/* webpackChunkName: "about" */ './views/EditReview.vue'),
      meta: { title:'Edit' }
    },
    {
      path: '/product/:id(\\d+)',
      name: 'productDetail',
      component: () => import(/* webpackChunkName: "about" */ './views/Product.vue'),
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
