<template>
  <div class="character">
    <div v-if="isAchieve" class="character_slide">
      <figure class="character_img character_img_illust03"></figure>
    </div>
    <div v-else-if="isLogin" class="character_slide">
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
    isLogin: function() {
      return this.$store.getters.isLogin
    },
    isAchieve: function() {
      return this.$store.getters.isAchieve
    }
  },
  watch: {
    isLogin: function(newVal) {
      return newVal;
    }
  },
  created() {
    // character.style.height = window.innerHeight + 'px';
    // display store data
    // console.log('hellow world vue:' + this.$store.state.count);
  },
  mounted() {
    const character = document.querySelector('.character');
    character.style.height = window.innerHeight + 'px';

    setTimeout(function() {
      document.querySelector('.character').classList.add('show');
    }, 1000)
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
.character_img_illust03 {
  background-image: url(../assets/img/eyevow/character_illust_03.png);
}
</style>
