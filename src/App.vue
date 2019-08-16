<template>
  <div id="app">
    <Header msg="This is the header! "/>
    <div id="contents">
      <router-view/>
    </div>
    <Nav />
  </div>
</template>

<style lang="scss">
@import '@/assets/style/style.scss';
</style>

<script>
import Header from '@/components/Header.vue'
import Nav from '@/components/Nav.vue'

export default {
  name: 'app',
  components: {
    Header,
    Nav
  },
  mounted () {
    const contents = document.querySelector('#contents');
    contents.style.height = window.innerHeight + 'px';
  }
}

</script>
