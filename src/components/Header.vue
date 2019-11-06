<template>
  <header id="header" :class="{'active': $route.path !== '/' || login }">
    <nav class="logo">
      <router-link to="/">
        <!-- <img alt="eyevow" src="@/assets/img/title_min.png" v-if="$route.path !== '/'"> -->
        <img alt="eyevow" src="@/assets/img/title.png">
      </router-link>
    </nav>
    <h1 class="title" v-if="$route.path !== '/'">{{ pageTitle }}</h1>
    <nav class="config">
      <router-link to="/edit">
        <!-- <span class="badge"></span> -->
        <figure class="icon">
          <img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png">
        </figure>
      </router-link>
    </nav>
  </header>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'Header',
  data() {
    return {};
  },
  computed: {
    ...mapGetters({
      pageTitle:'pageTitle',
    }),
    login() {
      return this.$store.state.app.isLogin
    }
  },
  watch: {
    login: function(newValue, oldValue) {
      console.log('newValue: ', newValue);
      console.log('oldValue: ', oldValue);
    }
  },
  created() {
    // console.log(this.$store.state.count);
    // console.log(location.pathname);
    console.log('login: ' + this.$store.state.app.isLogin);
  },
  mounted() {
    console.log('mounted: ',location.pathname);
  }
}

</script>

<style scoped>
</style>
