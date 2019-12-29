<template>
  <header id="header" :class="{'active': $route.path !== '/' || isVow }">
    <nav class="logo">
      <router-link to="/">
        <!-- <img alt="eyevow" src="@/assets/img/title_min.png" v-if="$route.path !== '/'"> -->
        <img alt="eyevow" src="@/assets/img/common/logo_i.svg">
      </router-link>
    </nav>
    <h1 class="title" v-if="$route.path !== '/'">{{ pageTitle }}</h1>
    <!--
    <nav class="config config_login" v-if="!isLogin">
      <a href="https://eyevow.work.suichu.net/id/">ログイン</a>
    </nav>
    <nav class="config" v-if="isLogin">
      <button @click="signout">
        <figure class="icon">
          <img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png">
        </figure>
      </button>
    </nav>
    -->
  </header>
</template>

<script>
// import { mapGetters } from 'vuex';

export default {
  name: 'Header',
  data() {
    return {};
  },
  methods: {
    signout: function() {
      this.$store.commit('signout');
      this.$router.push({ path: '/'})
    }
  },
  watch: {
    isLogin: function(newVal) {
      return newVal;
    },
    isVow: function(newVal) {
      return newVal;
    },
    isAchieve: function(newVal) {
      return newVal;
    }
  },
  computed: {
    pageTitle: function() {
      return this.$store.getters.pageTitle
    },
    isLogin: function() {
      return this.$store.getters.isLogin
    },
    isVow: function() {
      return this.$store.getters.isVow
    },
    isAchieve: function() {
      return this.$store.getters.isAchieve
    }
  },
  watch: {
    isLogin: function(newVal) {
      return newVal;
    }
  },
  created() {
    // console.log(this.$store.state.count);
    // console.log(location.pathname);
    // console.log('login: ' + this.$store.state.app.isLogin);
  },
  mounted() {
    // console.log('mounted: ',location.pathname);
  }
}

</script>

<style scoped>
</style>
