import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  strict: true,
  state: {
    count: 0,
    app: {
      login: false
    }
  },
  mutations: {
    // test:count up
    countup(state) {
      state.count ++
    },
    login(state) {
      state.app.login = true
    },
    logout(state) {
      state.app.login = false
    }
  },
  actions: {

  }
})
