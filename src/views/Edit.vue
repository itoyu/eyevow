<template>
  <div class="edit">
    <!-- <h1 class="page_title">Account</h1> -->
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">ここではアカウント情報の変更や、画像のダウンロードができるわ</span>
      </p>
    </div>

    <div class="edit_item">
      <h2 class="edit_tit">アカウント情報</h2>
      <div class="profile">
        <dl class="profile_name">
          <dt>Name</dt>
          <dd>
            <input type="text" name="profile_name" value="user name" />
            <button @click="updateProfile" :class="{'disabled': updateProfileReady}" class="btn btn_min">Save</button>
          </dd>
        </dl>
        <dl class="profile_photo">
          <dt>Photo</dt>
          <dd>
            <figure>
              <label for="file_upload" class="file_upload">
                <img alt="" src="@/assets/img/eyevow/icon_illust_01.png">
                <span class="cover"><img alt="" src="@/assets/img/icon_camera.svg"></span>
                <input type="file" id="file_upload" @change="updateProfilePhoto" />
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

export default {
  data() {
    return {
      edit: false
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
    updateProfilePhoto: function(e) {
      const file = e.srcElement.files[0];
      console.log(file);

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
    // this.$store.commit('countup')
  }
}
</script>
