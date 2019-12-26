<template>
  <div class="vow_complete" id="vow_complete">
    <h1 class="title">Complete !</h1>
    <dl class="vow_complete_item">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">おつかれさま。あなたの誓いを受けとったよ。</span>
        </p>
      </dt>
      <dd class="vow_complete_wrap">
        <div class="vow_complete_img" :class="myVow.type"></div>
        <div class="vow_complete_text">
          <p>{{myVow.text}}</p>
        </div>
      </dd>
    </dl>

    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">この誓いを残しておくために<br>アカウント登録してね。</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">SNSのアカウントで<br>アカウント登録 / ログインできるよ。</span>
      </p>

      <nav class="signup">
        <a href="https://eyevow.work.suichu.net/id/" class="signup_item signup_item_twitter btn"><iconArr />Twitter</a>
      </nav>
    </div>
  </div>
</template>

<script>
import iconArr from '@/assets/img/common/icon_arr_right.svg?inline'

export default {
  name: 'vowComplete',
  data: function() {
    return {
      myVow: {
        type: this.$store.getters.myVow.type,
        text: this.$store.getters.myVow.text,
        displayType: ('photo' === this.$store.getters.myVow.type) ? '写真' : 'イラスト'
      }
    }
  },
  components: {
    iconArr
  },
  created() {
    localStorage.setItem('vowType', this.$store.getters.myVow.type)
    localStorage.setItem('vowText', this.$store.getters.myVow.text)
  }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.vow_complete_img.illust {
  background-image: url(../assets/img/eyevow/character_illust_02.png);
}
.vow_complete_img.photo {
  background-image: url(../assets/img/eyevow/character_photo_02.png);
}
</style>
