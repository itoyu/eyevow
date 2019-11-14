<template>
  <div class="home">
    <div v-if="!login">
      <p class="catch">あなたの誓いを応援する相棒</p>
      <h1 class="title"><img alt="eyevow" src="@/assets/img/title.png"></h1>
    </div>
    <Character />
    <Chat />
    <Notice  v-if="login" />
    <!-- <StartBtn v-if="!login" /> -->
  </div>
</template>

<script>
// @ is an alias to /src
import Character from '@/components/Character.vue'
import Chat from '@/components/Chat.vue'
import Notice from '@/components/Notice.vue'
// import StartBtn from '@/components/StartBtn.vue'

export default {
  name: 'Home',
  components: {
    Character,
    Chat,
    Notice,
    // StartBtn
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
    }
  },
  computed: {
    login: function() {
      return this.$store.state.app.isLogin
    }
  },
  created() {
    this.init();
  }
}
</script>

<style lang="scss" scoped>
#contents {
  height: 100vh;
}
</style>
