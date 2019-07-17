<template>
  <header class="global">
    <h1>{{ msg }}</h1>
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

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  header.global {
    text-align: center;
    background: #ccc;
  }
</style>
