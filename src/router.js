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
      meta: { title: 'Home' },
      component: Home
    },
    {
      path: '/about/:id',
      meta: { title: 'About' },
      component: () => import('./views/About.vue'),

      children: [
        {
          path: '/about/',
          meta: { title: 'About' },
          component:() => import('./views/AboutTutorial')
        },
        {
          path: '/about/statement',
          meta: { title: 'Statement' },
          component:() => import('./views/AboutStatement')
        },
        {
          path: '/about/credit',
          meta: { title: 'Credit' },
          component:() => import('./views/AboutCredit')
        },
        {
          path: '/about/story',
          meta: { title: 'Story' },
          component:() => import('./views/AboutStory')
        },
        {
          path: '/about/story/:id', //:id(\\d+)
          props: route => ({ 'id': Number(route.params.id) }),
          meta: route => ({ title: `Story ${route.params.id}` }),
          component: () => import('./views/AboutStoryDetail.vue')
        }
      ]
    },
    {
      path: '/vow',
      meta: { title: 'Vow' },
      component: () => import('./views/Vow.vue')
    },
    {
      path: '/achieve',
      meta: { title: 'Achieved' },
      component: () => import('./views/Achieve.vue')
    },
    {
      path: '/achieved',
      meta: { title: 'Achieved' },
      component: () => import('./views/Achieved.vue')
    },
    {
      path: '/eyevow',
      meta: { title: 'eyevow' },
      component: () => import('./views/Eyevow.vue')
    },
    {
      path: '/cheer',
      meta: { title: 'Cheer' },
      component: () => import('./views/Cheer.vue')
    },
    {
      path: '/edit',
      meta: { title: 'Edit' },
      component: () => import('./views/Edit.vue')
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
