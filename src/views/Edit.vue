<template>
  <div class="edit">
    <!-- <h1 class="page_title">Account</h1> -->
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">ここではアカウント情報の変更や、画像のダウンロードができるよ</span>
      </p>
    </div>

    <div class="edit_item">
      <h2 class="edit_tit">アカウント情報</h2>
      <div class="profile">
        <dl class="profile_name">
          <dt>Name</dt>
          <dd>
            <input type="text" name="profile_name" :value="userInfo.name" />
            <button @click="updateProfile" :class="{'disabled': updateProfileReady}" class="btn btn_min">Save</button>
          </dd>
        </dl>
        <dl class="profile_photo">
          <dt>Photo</dt>
          <dd>
            <figure>
              <label for="file_upload" class="file_upload">
                <img alt="" :src="userInfo.icon">
                <span class="cover"><img alt="" src="@/assets/img/common/icon_camera.svg"></span>
                <input type="file" id="file_upload" accept="image/jpeg, image/png" @change="updateProfilePhoto" />
              </label>
            </figure>
          </dd>
        </dl>
      </div>
    </div>

    <div class="edit_item">
      <h2 class="edit_tit">あなたの誓い</h2>
      <div class="myvow">
        <dl class="myvow_txt">
          <dt>Vow</dt>
          <dd>たくさん勉強して宇宙の星々を研究する学者になるぞ！たくさん勉強して宇宙の星々を研究する学者になるぞ！</dd>
        </dl>
        <dl class="myvow_type">
          <dt>Type</dt>
          <dd>
            <figure>
              <img alt="" src="@/assets/img/eyevow/icon_illust_01.png">
            </figure>
          </dd>
        </dl>
      </div>
      <div class="preview">
        <button @click="openPopup" class="btn btn_normal">View Card</button>
      </div>
    </div>

    <div class="popup popup_confirm">
      <div class="popup_in">
        <button @click="closePopup" class="btn btn_closePopup">×</button>
        <div class="vow_complete_wrap">
          <div class="vow_complete_img">
            <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_02.png"></figure>
          </div>
          <div class="vow_complete_text">
            <p>たくさん勉強して宇宙の星々を研究する学者になるぞ！</p>
          </div>
          <span class="effect effect01"><img alt="" src="@/assets/img/effect_card.gif"></span>
          <span class="effect effect02"><img alt="" src="@/assets/img/effect_card.gif"></span>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import api from '@/api/client';

export default {
  data() {
    return {
      edit: false,
      userInfo: {
        // name: this.$store.getters.userInfo.name,
        // icon: this.$store.getters.userInfo.icon
        name: this.$store.state.devData.user.name,
        icon: this.$store.state.devData.user.icon
      }
    }
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
      document.querySelector('body').classList.remove('focus_vow');
    },
    updateProfile: function() {
      console.log('updateProfile');
    },
    // プロフィール画像の変更
    getBase64 (file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = () => resolve(reader.result)
        reader.onerror = error => reject(error)
      })
    },
    updateProfilePhoto: function(e) {
      const images = e.target.files || e.dataTransfer.files
      this.getBase64(images[0])
        // .then(image => this.avatar = image)
        .then(image => {
          var dataUrl = image;
          this.userInfo.icon = dataUrl;
          // console.log(dataUrl);
          // this.vowlist = json.vows.reverse();

          api.patch('/user', {
            name: 'guest',
            icon: dataUrl
          }, {
            headers: {
              Authorization: `Bearer ${this.$store.state.devData.token}`,
            }
          })
            // .then(function() {
            //   console.log('update user info');
            // })
            // .catch(function (error) {
            //   // handle error
            //   console.log(error);
            // });
        })
        .catch(error => this.setError(error, 'error upload image.'))
    },
    openPopup: function() {
      document.querySelector('.popup').classList.add('show');
    },
    closePopup: function() {
      document.querySelector('.popup').classList.remove('show');
    }
  },
  computed: {
    updateProfileReady() {
      return (this.edit) ? false : true;
    }
  },
  created() {
    this.init();
  }
}
</script>
