import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

const store = new Vuex.Store({
  strict: true,
  state: {
    devData: {
      token:  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiNWRmNGYzZDk2YTM0MzMxNGNlODUxNDM2IiwiZXhwIjoxNzM0MDE0Mjk3fQ.37yVyjK5fR9JVc3MOgqZUkpmDJlLTDQ61gPSWFIs1-o',
      user: {
        isLogin: false,
        isAchieve: false,
        isVow: false,
        id: '5df4f3d96a343314ce851436',
        name: 'ゲスト',
        icon: 'http://eyevow.work.suichu.net/blob/user/NUKwnnfWg.jpg'
      }
    },
    app: {
      pageTitle: 'Home'
    },
    user: {
      isLogin: false,
      isAchieve: false,
      isVow: false,
      id: String,
      name: String,
      icon: String
    },
    myVows: [
      // {
      //   id: String,
      //   type: String,
      //   text: String,
      //   cheer_count: Number,
      //   archived: false
      // }
    ],
    allVows: [
      // {
      //   id: String,
      //   user: {
      //     id: String,
      //     name: String,
      //     icon: String
      //   },
      //   text: String,
      //   cheer_count: Number,
      //   archived: false
      // }
    ]
  },
  getters: {
    pageTitle(state) {
      return state.app.pageTitle;
    },
    isLogin(state) {
      return state.user.isLogin;
    }
  },
  actions: {
    changePage({ commit }, title) {
      commit('cahgePageTitle', title);
    },
  },
  mutations: {
    cahgePageTitle(state, title) {
      state.app.pageTitle = title;
    },
    cheerCountup(state) {
      // state.count ++
    },
    isLogin(state) {
      state.user.isLogin = true;
    },
    isLogout(state) {
      state.user.isLogin = false;
    },
    setVow(state) {
      state.user.isVow = true;
    },
    unsetVow(state) {
      state.user.isVow = false;
    },
    setAchieve(state) {
      state.user.isAchieve = true;
    },
    unsetAchieve(state) {
      state.user.isAchieve = false;
    }
  },
  plugins: [createPersistedState()]
})
export default store;
