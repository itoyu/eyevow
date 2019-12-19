<template>
  <div class="vow_complete" id="vow_complete">
    <h1 class="title">Complete !</h1>
    <dl class="vow_complete_item">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">おつかれさま。あなたの誓い、受けとったわ。</span>
        </p>
      </dt>
      <dd class="vow_complete_wrap">
        <div class="vow_complete_img">
          <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_02.png"></figure>
        </div>
        <div class="vow_complete_text">
          <p>たくさん勉強して宇宙の星々を研究する学者になるぞ！</p>
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
        <button @click="signup('twitter')" class="signup_item signup_item_twitter btn"><iconArr />Twitter</button>
        <button @click="signup('line')" class="signup_item signup_item_line btn"><iconArr />Line</button>
        <button @click="signup('google')" class="signup_item signup_item_google btn"><iconArr />Google</button>
      </nav>
    </div>

    <div class="popup popup_signup">
      <div class="popup_in">
        <div style="text-align: center;">
          <p style="margin-bottom: 20px;">Social login...</p>
          <button class="btn btn_vow" @click="completeSignup">Home</button>
        </div>
      </div>
    </div>



  </div>
</template>

<script>
import api from '@/api/client';
import iconArr from '@/assets/img/icon_arr_right.svg?inline'

export default {
  name: 'vowComplete',
  data: function() {
    return {
      testUser: {
        token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiNWRmNGYzZDk2YTM0MzMxNGNlODUxNDM2IiwiZXhwIjoxNzM0MDE0Mjk3fQ.37yVyjK5fR9JVc3MOgqZUkpmDJlLTDQ61gPSWFIs1-o',
        id: '5df4f3d96a343314ce851436',
      }

    }
  },
  components: {
    iconArr
  },
  created() {
    // console.log('setVow');
    this.$store.commit('setVow');
  },
  methods: {
    signup: function(s) {
      console.log(s);
      document.querySelector('.popup').classList.add('show');
    },
    completeSignup: function() {



      // console.log('completeSignup');
      this.$store.commit('isLogin');
      // window.location = "/"

      // window.location = '/';
      // console.log(this.$store.state.app.isLogin);
      setTimeout(function() {
        window.location = '/';
      },500)
    }
  },
  mounted () {
    // #Set Vow
    api.post('/vows', { 'text':'誓いテキスト' }, {
      headers: {
        Authorization: `Bearer ${this.testUser.token}`,
      }
    })
      .then(res => res.data)
      .then(json => {
        console.log('set vow');
        console.log(json);
      })

    // #Get User
    // api.get('/user', {
    //   headers: {
    //     Authorization: `Bearer ${this.token}`,
    //   }
    // })
    //   .then(res => res.data)
    //   .then(json => {
    //     console.log('user list');
    //     console.log(json);
    //   })
  }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
