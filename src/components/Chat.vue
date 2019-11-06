<template>

  <div class="chat" :class="{'member': login }">
    <div class="chat_wrap" v-if="!login">
      <p v-for="item in guestChats" v-bind:key="item.id" class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt" id="`${ item.id }`">{{ item.text }}</span>
      </p>
    </div>
    <div class="chat_wrap" v-if="login">
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
  watch: {
    login: function(newValue) {
      return newValue;
    }
  },
  methods: {
    init: function() {
      const chatWrap = document.querySelector('.chat'),
            chatItems = document.querySelectorAll('.chat_item'),
            isLogin = this.$store.state.app.isLogin;
      var initChatSet, chatCnt = 0;

      console.log('isLogin: ' + isLogin);

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
        if(chatCnt >= 4 && !isLogin) {
          chatStop();
        }
        if(chatCnt >= 3 && isLogin) {
          chatStop();
        }
        if(chatCnt === 4 && !isLogin) {
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
  computed: {
    login: function() {
      return this.$store.state.app.isLogin
    }
    // homeMessage: function() {
    //   // chatMessage.fetch();
    //   console.log();
    // }
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
