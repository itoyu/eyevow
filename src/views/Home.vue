<template>
  <div class="home">
    <h1 class="title"><img alt="eyevow" src="@/assets/img/title.png"></h1>
    <Character />
    <Notice />
    <StartBtn />
  </div>
</template>

<script>
// @ is an alias to /src
import Character from '@/components/Character.vue'
import Notice from '@/components/Notice.vue'
import StartBtn from '@/components/StartBtn.vue'

export default {
  name: 'home',
  components: {
    Character,
    Notice,
    StartBtn
  }
}
</script>

<style lang="scss" _scoped>
#header {
  // display: none;
}
.chat {

}
</style>
