<template>

  <div class="chat" :class="{'member': isVow }">
    <div class="chat_wrap" v-if="!isVow">
      <p v-for="item in guestChats" v-bind:key="item.id" class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt" id="`${ item.id }`">{{ item.text }}</span>
      </p>
    </div>
    <div class="chat_wrap" v-if="isVow">
      <p v-for="item in memberChats" v-bind:key="item.id" class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_02.png"></span>
        <span class="txt" id="`${ item.id }`" v-html="`${ item.text }`"></span>
      </p>
    </div>
  </div>
</template>

<script>
import guestMessage from '@/api/guestMessage.js'
import memberMessage from '@/api/memberMessage.js'

export default {
  name: 'Chat',
  data: function() {
    return {
      guestChats: guestMessage.fetch(),
      memberChats: memberMessage.fetch(),
      displayMessageNum: 3
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
  methods: {
    init: function() {
      const chatWrap = document.querySelector('.chat'),
            chatItems = document.querySelectorAll('.chat_item'),
            stateLogin = this.$store.state.user.isLogin;
      var initChatSet, chatCnt = 0;

      function chatStart() {
        initChatSet = setInterval(chatPush, 3000);
      }
      function chatStop() {
        clearInterval(initChatSet);
      }
      function chatPush() {
        chatItems[chatCnt].classList.add('show');

        // #上限3件表示にする
        // if(chatCnt > 2) {
        //   chatItems[chatCnt - 3].classList.add('hide');
        // }

        chatCnt ++;

        // if(chatCnt >= chatItems.length) {
        if(chatCnt >= 4 && !stateLogin) {
          chatStop();
        }
        if(chatCnt >= 3 && stateLogin) {
          chatStop();
        }
        if(chatCnt === 4 && !stateLogin) {
          if(document.getElementById('startbtn')) {
            document.getElementById('startbtn').classList.add('show');
          }
        }
      }

      chatWrap.onclick = function(e) {
        e.preventDefault();
        chatStop();
        if(chatCnt < chatItems.length) {
          chatPush();
        }
      }

      // chatPush();
      chatStart();
    }
  },
  mounted() {
    this.init();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .chat {
    &_item {
      display: none;
    }
  }
</style>
