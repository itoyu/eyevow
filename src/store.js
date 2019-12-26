import Vue from 'vue'
import Vuex from 'vuex'
import createPersistedState from 'vuex-persistedstate'

Vue.use(Vuex)

const store = new Vuex.Store({
  strict: true,
  state: {
    devData: {
      token:  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiNWUwNDJiNzllZGJmYjc3MDVkODJjZjc1IiwiZXhwIjoxNzM1MDExNTc3fQ.QYgFEQi09Uq3ZqbQLnuFKXcU19AUGDj_sR_1B7am7aw',
      user: {
        isLogin: false,
        isVow: false,
        isAchieve: false,
        id: '5e01bd46edbfb7705d82cf5e',
        name: 'ゲスト',
        icon: 'http://eyevow.work.suichu.net/blob/user/qu63MqBZR.jpg'
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
      icon: '',
      token: ''
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
    userInfo(state) {
      return state.user;
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
    setUserToken({ commit }, token) {
      commit('setToken', token);
    },
    putTemporaryVow({ commit }, {vowType, vowText }) {
      const putVow = { type: vowType, text: vowText }
      commit('changeMyVow', putVow);
    },
    initSetVow({ commit }, {vowType, vowText }) {
      const setVow = { type: vowType, text: vowText }
      commit('setMyVow', setVow);
    }
  },
  mutations: {
    cahgePageTitle(state, title) {
      state.app.pageTitle = title;
    },
    setToken(state, token) {
      state.user.token = token;
    },
    changeMyVow(state, putVow) {
      state.myVow.type = putVow.type;
      state.myVow.text = putVow.text;
    },
    setMyVow(state, setVow) {
      state.user.isLogin = true;
      state.user.isVow = true;
      state.myVow.type = setVow.type;
      state.myVow.text = setVow.text;
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
    },
    signout(state) {
      state.user.token = '';
      state.user.isLogin = false;
      state.user.isVow = false;
      state.user.isAchieve = false;
      state.myVow.id = '';
    }
  },
  plugins: [createPersistedState()]
})
export default store;
