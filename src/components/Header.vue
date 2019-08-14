<template>
  <header id="header">
    <nav class="logo">
      <router-link to="/">
        <img alt="eyevow" src="@/assets/img/logo.png">
      </router-link>
    </nav>
    <p class="message">{{ msg }}</p>
    <nav class="config">
      <router-link to="/edit">
        <span class="notice"></span>
        <figure class="icon">
          <img alt="profile icon" src="@/assets/img/icon_eyevow_illust.png">
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
  }
}
</script>

<style scoped>
  /* header.global {
    text-align: center;
    background: #ccc;
  } */
</style>
