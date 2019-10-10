<template>
  <div class="popup" id="vow_popup">
    <div class="popup_in">
      <div class="popup_message">
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">新しいeyevowを追加しますか？</span>
        </p>
      </div>
      <nav class="popup_nav">
        <router-link to="/" class="btn btn_back">
          <img alt="Back" src="@/assets/img/icon_arr_left.svg">
        </router-link>
        <router-link to="#" class="btn btn_normal">+ NEW VOW</router-link>
      </nav>
    </div>
  </div>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
