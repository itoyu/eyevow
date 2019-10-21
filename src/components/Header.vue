<template>
  <header id="header" v-if="login">
    <nav class="logo">
      <router-link to="/">
        <img alt="eyevow" src="@/assets/img/logo.png">
      </router-link>
    </nav>
    <p class="message">{{ msg }}</p>
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

export default {
  name: 'Header',
  props: {
    msg: String
  },
  computed: {
    login() {
      return this.$store.state.app.login
    }
  },
  created() {
    // console.log(this.$store.state.count);
  }
}

</script>

<style scoped>
  /* header.global {
    text-align: center;
    background: #ccc;
  } */
</style>
