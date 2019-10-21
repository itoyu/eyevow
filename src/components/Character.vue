<template>
  <div class="character">
    <div v-if="login">
      <p>login user</p>
      <figure class="character_img character_img_illust01"></figure>
    </div>
    <div v-else>
      <figure class="character_img character_img_illust01"></figure>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Character',
  computed: {
    login() {
      return this.$store.state.app.login
    }
  },
  created() {
    // character.style.height = window.innerHeight + 'px';
    // display store data
    // console.log('hellow world vue:' + this.$store.state.count);
  },
  mounted () {
    const character = document.querySelector('.character_img');
    character.style.height = window.innerHeight + 'px';
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
.character_img_illust01 {
  background-image: url(../assets/img/eyevow/character_illust_01.png);
}
</style>
