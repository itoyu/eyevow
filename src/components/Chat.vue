<template>

  <div class="chat">
    <div class="chat_wrap">
      <p v-for="item in chats" v-bind:key="item.id" class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt" id="`${ item.id }`">{{ item.text }}</span>
      </p>
    </div>
  </div>
</template>

<script>
import chatMessage from '@/api/chatMessage.js'

export default {
  name: 'Chat',
  data: function() {
    return {
      chats: chatMessage.fetch(),
      displayMessageNum: 3
    }
  },
  methods: {
    init: function() {
      const chatWrap = document.querySelector('.chat'),
            chatItems = document.querySelectorAll('.chat_item');
      var initChatSet, chatCnt = 0;

      function chatStart() {
        initChatSet = setInterval(chatPush, 4000);
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
        if(chatCnt >= 4) {
          chatStop();
        }
        if(chatCnt === 4) {
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

      chatStart();
    }
  },
  computed: {
    // login: function() {
    //   return this.$store.state.app.login
    // }
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
</style>
