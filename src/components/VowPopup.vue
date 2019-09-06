<template>
  <div class="popup" id="vow_popup">
    <div class="popup_in">
      <p class="popup_message">新しいeyevowを追加しますか？</p>
      <nav class="popup_nav">
        <router-link to="/" class="btn btn_back">
          <img alt="Back" src="@/assets/img/icon_arr_left.svg">
        </router-link>
        <router-link to="/edit/review" class="btn btn_normal">NEW VOW</router-link>
      </nav>
    </div>
  </div>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
