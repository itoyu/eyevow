import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import VuePageTransition from 'vue-page-transition'
import ToggleButton from 'vue-js-toggle-button'
Vue.use(VuePageTransition)
Vue.use(ToggleButton)

// console message
Vue.config.productionTip = true

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')