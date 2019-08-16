<template>
  <div class="edit">
    <!-- <h1 class="page_title">Account</h1> -->

    <h2 class="edit_tit">Account</h2>
    <div class="profile">
      <dl class="profile_name">
        <dt>Name</dt>
        <dd><input type="text" name="profile_name" value="user name" /></dd>
      </dl>
      <dl class="profile_photo">
        <dt>Photo</dt>
        <dd>
          <router-link to="/edit/photo" class="profile_img_edit">
            <figure>
              <img alt="" src="@/assets/img/eyevow/icon_illust_01.png">
              <span class="cover"><img alt="" src="@/assets/img/icon_camera.svg"></span>
            </figure>
          </router-link>
        </dd>
      </dl>
    </div>

    <h2 class="edit_tit">My vow</h2>
    <div class="myvow">
      <dl class="myvow_txt">
        <dt>Vow</dt>
        <dd>たくさん勉強して宇宙の星々を研究する学者になるぞ！</dd>
      </dl>
      <dl class="myvow_type">
        <dt>Type</dt>
        <dd>
          <figure>
            <img alt="" src="@/assets/img/eyevow/icon_illust_01.png">
          </figure>
          <p>Illust</p>
        </dd>
      </dl>
      <dl class="myvow_preview">
        <dt>
          <p class="txt">eyevowの画像が確認できます。</p>
          <router-link to="/edit/photo" class="btn btn_noram">View</router-link>
        </dt>
        <dd>
          <p class="txt">Share</p>
          <a href="#" class="btn btn_sns" target="_blank">
            <img alt="Share to Twitter" src="@/assets/img/icon_twitter.svg">
          </a>
          <a href="#" class="btn btn_sns" target="_blank">
            <img alt="Share to Line" src="@/assets/img/icon_line.svg">
          </a>
        </dd>
      </dl>
    </div>
  </div>

</template>

<script>

export default {
  created() {
    // display store data
    // console.log('edit vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
