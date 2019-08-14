<template>
  <div class="character">
    <figure class="character_img"></figure>

    <div class="chat">
      <p class="chat_item">誓いを込めてくれてありがとう</p>
      <p class="chat_item">誓い達成に向けてがんばってね！</p>
      <p class="chat_item">誓い達成に向けてがんばってね！誓い達成に向けてがんばってね！誓い達成に向けてがんばってね！</p>
    </div>

  </div>
</template>

<script>
export default {
  name: 'Character',
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
.character_img {
  background-image: url(../assets/img/eyevow/character_illust_01.png);
}
</style>
