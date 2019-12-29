<template>
  <div class="home">
    <!--
    <nav class="home_login" v-if="!isLogin">
      <a href="https://eyevow.work.suichu.net/id/">ログイン</a>
    </nav>
    -->
    <div v-if="!isVow">
      <p class="catch">あなたの誓いを応援する相棒</p>
      <h1 class="title"><img alt="eyevow" src="@/assets/img/common/logo.svg"></h1>
    </div>
    <Character />
    <Chat />
    <Notice v-if="isVow" />
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
      document.querySelector('body').classList.remove('focus', 'focus_vow');
    }
  },
  watch: {
    isLogin: function(newVal) {
      return newVal;
    },
    isVow: function(newVal) {
      return newVal;
    },
    isAchieve: function(newVal) {
      return newVal;
    }
  },
  computed: {
    isLogin: function() {
      return this.$store.getters.isLogin
    },
    isVow: function() {
      return this.$store.getters.isVow
    },
    isAchieve: function() {
      return this.$store.getters.isAchieve
    }
  },
  created() {
    this.init();
  },
  mounted () {
    const contents = document.querySelector('#contents .home');
    contents.style.height = (window.innerHeight - 150) + 'px';

    window.addEventListener('resize', function() {
      contents.style.height = (window.innerHeight - 150) + 'px';
    });
  }
}
</script>
