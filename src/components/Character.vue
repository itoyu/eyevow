<template>
  <div class="character">
    <div v-if="login" class="character_slide">
      <p>login user</p>
      <figure class="character_img character_img_illust02"></figure>
    </div>
    <div v-else class="character_slide">
      <figure class="character_img character_img_illust01"></figure>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Character',
  computed: {
    login() {
      return this.$store.state.app.isLogin
    }
  },
  watch: {
    login: function(newValue) {
      return newValue;
    }
  },
  created() {
    // character.style.height = window.innerHeight + 'px';
    // display store data
    // console.log('hellow world vue:' + this.$store.state.count);
  },
  mounted () {
    const character = document.querySelector('.character');
    character.style.height = window.innerHeight + 'px';
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
.character_img_illust01 {
  background-image: url(../assets/img/eyevow/character_illust_01.png);
}
.character_img_illust02 {
  background-image: url(../assets/img/eyevow/character_illust_02.png);
}
</style>
