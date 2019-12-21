<template>
  <header id="header" :class="{'active': $route.path !== '/' || isLogin }">
    <nav class="logo">
      <router-link to="/">
        <!-- <img alt="eyevow" src="@/assets/img/title_min.png" v-if="$route.path !== '/'"> -->
        <img alt="eyevow" src="@/assets/img/logo_i.svg">
      </router-link>
    </nav>
    <h1 class="title" v-if="$route.path !== '/'">{{ pageTitle }}</h1>

    <!-- <nav class="config">
      <button @click="signout">
        <figure class="icon">
          <img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png">
        </figure>
      </button>
    </nav> -->
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
      this.$store.commit('isLogout');
      this.$store.commit('unsetAchieve');
    }
  },
  computed: {
    pageTitle: function() {
      return this.$store.getters.pageTitle
    },
    isLogin: function() {
      return this.$store.getters.isLogin
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
