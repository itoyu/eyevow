<template>
  <div class="cheer">
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">みんなの誓いがここに集まっているわ</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">応援したいものがあったら「Cheer」ボタンを押して声援を送ってあげて</span>
      </p>
    </div>

    <ul class="tab_nav">
      <li class="tab_nav_item active">Hot</li>
      <li class="tab_nav_item">Progress</li>
      <li class="tab_nav_item">Achived</li>
    </ul>

    <div class="tab_contents">
      <div class="tab_contents_item">

        <ul class="cheer_list">
          <li class="cheer_list_item">
            <figure class="cheer_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
            <div class="cheer_info">
              <p class="cheer_name">Name</p>
              <p class="cheer_text">エモすぎる楽曲を作って世界一のロックバンドを結成する！</p>
              <p class="cheer_time">2h</p>
            </div>
            <div class="cheer_action current">
              <p class="cheer_action_num">3</p>
              <button class="cheer_action_btn"><iconLike /></button>
            </div>
          </li>
          <li class="cheer_list_item">
            <figure class="cheer_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
            <div class="cheer_info">
              <p class="cheer_name">User Name User Name User Name User Name User Name </p>
              <p class="cheer_text">エモすぎる楽曲を作って世界一のロックバンドを結成する！</p>
              <p class="cheer_time">19.12.27</p>
            </div>
            <div class="cheer_action">
              <!-- <p class="cheer_action_num">3</p> -->
              <button class="cheer_action_btn"><iconLike /></button>
            </div>
          </li>
          <li class="cheer_list_item achived">
            <span class="icon_achieve"><iconAchieve /></span>
            <figure class="cheer_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
            <div class="cheer_info">
              <p class="cheer_name">Name</p>
              <p class="cheer_text">たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。</p>
              <p class="cheer_time">19.11.03</p>
            </div>
            <div class="cheer_action">
              <!-- <p class="cheer_action_num">3</p> -->
              <button class="cheer_action_btn"><iconLike /></button>
            </div>
          </li>
          <li class="cheer_list_item achived">
            <span class="icon_achieve"><iconAchieve /></span>
            <figure class="cheer_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
            <div class="cheer_info">
              <p class="cheer_name">Name</p>
              <p class="cheer_text">たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。たくさん勉強して宇宙の星々を研究する学者になります。</p>
              <p class="cheer_time">19.10.03</p>
            </div>
            <div class="cheer_action">
              <p class="cheer_action_num">333</p>
              <button class="cheer_action_btn"><iconLike /></button>
            </div>
          </li>
        </ul>
      </div>
    </div>

    <!-- <ul>
      <li v-for="(item, index) in list" v-bind:key="item.id">
        ID:{{item.id}}
        User:{{item.user.id}}
        Vow:{{item.text}}
        Cheer:{{item.cheer_count}}
        Achieved:{{item.archived}}
        Index: {{index}}
      </li>
    </ul> -->
  </div>
</template>

<script>
// import axios from 'axios'
import vows from '@/api/vows.json'
import iconLike from '@/assets/img/icon_like.svg?inline'
import iconAchieve from '@/assets/img/icon_achieve.svg?inline'

export default {
  name: 'Cheer',
  data: function() {
    return {
      list: []
    }
  },
  components: {
    iconLike,
    iconAchieve
  },
  created() {
    // axios
    //   .get('./vows.json')
    //   .then(function(res) {
    //     this.list = res.data
    // })
    this.list = vows;
    // console.log(this.list);
  }
}
</script>
