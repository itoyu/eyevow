import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  strict: true,
  state: {
    count: 0,
    pageTitle: 'Home',
    app: {
      login: false
    }
  },
  actions: {
    changePage({ commit }, title) {
      commit('cahgePageTitle', title);
    },
  },
  mutations: {
    // test:count up
    cahgePageTitle(state, title) {
      const st = state;
      st.pageTitle = title;
    },
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
  getters: {
    pageTitle(state) {
      const st = state;
      return st.pageTitle;
    },
  },
})
export default store;
