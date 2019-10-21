<template>
  <header id="header" :class="{'active': $route.path !== '/' || login }">
    <nav class="logo">
      <router-link to="/">
        <img alt="eyevow" src="@/assets/img/logo.png">
      </router-link>
    </nav>
    <p class="message">title: {{ pageTitle }}</p>
    <nav class="config">
      <router-link to="/edit">
        <span class="badge"></span>
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
      return this.$store.state.app.login
    }
  },
  created() {
    // console.log(this.$store.state.count);
    // console.log(location.pathname);
  },
  mounted() {
    console.log('mounted: ',location.pathname);
  }
}

</script>

<style scoped>
</style>
