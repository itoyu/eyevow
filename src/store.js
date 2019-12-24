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
        isVow: false,
        isAchieve: false,
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
      isVow: false,
      isAchieve: false,
      id: '',
      name: '',
      icon: ''
    },
    myVow: {
      id: '',
      type: 'illust',
      text: '',
      cheer_count: 0,
      set: false,
      archived: false
    },
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
    myVow(state) {
      return state.myVow;
    },
    isLogin(state) {
      return state.user.isLogin;
    },
    isVow(state) {
      return state.user.isVow;
    },
    isAchieve(state) {
      return state.user.isAchieve;
    }
  },
  actions: {
    changePage({ commit }, title) {
      commit('cahgePageTitle', title);
    },
    putTemporaryVow({ commit }, {vowType, vowText }) {
      const putVow = { type: vowType, text: vowText }
      commit('changeMyVow', putVow);
    }
  },
  mutations: {
    cahgePageTitle(state, title) {
      state.app.pageTitle = title;
    },
    changeMyVow(state, putVow) {
      state.myVow.type = putVow.type;
      state.myVow.text = putVow.text;
    },
    cheerCountup() {
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
  // plugins: [createPersistedState()]
})
export default store;
