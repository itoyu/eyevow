import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

const store = new Vuex.Store({
  strict: true,
  state: {
    count: 0,
    pageTitle: 'Home',
    vow: false,
    achieve: false,
    app: {
      isAchieve: false,
      isLogin: false,
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
    isLogin(state) {
      state.app.isLogin = true;
    },
    isLogout(state) {
      state.app.isLogin = false;
    },
    setVow(state) {
      state.vow = true;
    },
    unsetVow(state) {
      state.vow = false;
    },
    setAchieve(state) {
      state.app.isAchieve = true;
    },
    unsetAchieve(state) {
      state.app.isAchieve = false;
    }
  },
  getters: {
    pageTitle(state) {
      const st = state;
      return st.pageTitle;
    },
  },
  plugins: [createPersistedState()]
})
export default store;
