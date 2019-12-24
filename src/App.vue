<template>
  <div id="app">
    <Header msg="This is the header! "/>
    <div id="contents">
      <vue-page-transition name="fade">
        <router-view/>
      </vue-page-transition>
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
import api from './api/client';

export default {
  name: 'app',
  components: {
    Header,
    Nav
  },
  mounted () {
    api.get('/vows')
      .then(res => res.data)
      // .then(() => console.log("API is success"))
      .then(json => console.log(json));
    //
    // // #Get User
    // api.get('/user', {
    //   headers: {
    //     Authorization: `Bearer ${this.$store.state.devData.token}`,
    //   }
    // })
    //   .then(res => res.data)
    //   .then(json => {
    //     console.log('User list');
    //     console.log(json);
    //   })

    const contents = document.querySelector('#contents');
    contents.style.height = window.innerHeight + 'px';

    window.addEventListener('resize', function() {
      contents.style.height = window.innerHeight + 'px';
    });
  }
}

</script>
