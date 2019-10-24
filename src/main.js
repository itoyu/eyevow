import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ToggleButton from 'vue-js-toggle-button'
Vue.use(ToggleButton)

// console message
Vue.config.productionTip = true

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
