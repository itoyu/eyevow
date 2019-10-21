<template>
  <div class="home">
    <p class="catch">あなたの誓いを応援する相棒</p>
    <h1 class="title"><img alt="eyevow" src="@/assets/img/title.png"></h1>
    <Character />
    <Chat />
    <Notice  v-if="login" />
    <StartBtn  v-if="!login" />
  </div>
</template>

<script>
// @ is an alias to /src
import Character from '@/components/Character.vue'
import Chat from '@/components/Chat.vue'
import Notice from '@/components/Notice.vue'
import StartBtn from '@/components/StartBtn.vue'

export default {
  name: 'Home',
  components: {
    Character,
    Chat,
    Notice,
    StartBtn
  },
  computed: {
    login: function() {
      return this.$store.state.app.login
    }
  }
}
</script>

<style lang="scss" _scoped>
#header {
  // display: none;
}
.chat {
  margin-bottom: 70px;
}
</style>
