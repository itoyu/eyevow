<template>
  <div class="vow_complete" id="vow_complete">
    <h1 class="title">Congratulation !</h1>
    <dl class="vow_complete_item">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_03.png"></span>
          <span class="txt">誓い達成おめでとう！</span>
        </p>
      </dt>
      <dd class="vow_complete_wrap">
        <div class="vow_complete_img">
          <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_03.png"></figure>
        </div>
        <div class="vow_complete_text">
          <p>たくさん勉強して宇宙の星々を研究する学者になるぞ！</p>
        </div>
      </dd>
    </dl>

    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_03.png"></span>
        <span class="txt">誓い達成の特典があるわよ</span>
      </p>
    </div>

    <nav class="reward">
      <h2 class="title">REWARDS</h2>
      <router-link to="/about/story/epilogue" class="nav_item reward_item btn"><iconArr />EPILOGUE</router-link>
      <a href="https://eyevow.stores.jp/" target="_blank" class="nav_item reward_item btn"><iconArr />ONLINE SHOP</a>
      <router-link to="/" class="nav_item btn btn_back"><iconArr />HOME</router-link>
    </nav>

  </div>
</template>

<script>
import iconArr from '@/assets/img/common/icon_arr_right.svg?inline'

export default {
  name: 'AchieveComplete',
  components: {
    iconArr
  },
  methods: {
    init: function() {
      console.log('Achieved');
      this.$store.commit('setAchieve');
    }
  },
  created() {
    this.init();
  }
}
</script>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
