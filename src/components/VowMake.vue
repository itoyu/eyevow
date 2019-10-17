<template>
  <div class="vow_make" id="vow_make">
    <!-- Step 01 -->
    <dl class="vow_make_item step01">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">【Step 01】eyevowのタイプを選んでください。</span>
        </p>
      </dt>
      <dd>
        <div class="vow_make_type">
          <swiper :options="swiperOption">
            <swiper-slide>
              <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_01.png"></figure>
              <p class="txt">イラスト</p>
            </swiper-slide>
            <swiper-slide>
              <figure class="img"><img alt="" src="@/assets/img/eyevow/character_cosplayers_02.png"></figure>
              <p class="txt">レイヤー</p>
            </swiper-slide>
            <swiper-slide>
              <figure class="img soon">Coming Soon.</figure>
              <p class="txt">フィギュア</p>
            </swiper-slide>
          </swiper>
          <div class="swiper-pagination"></div>
        </div>
        <nav class="vow_make_nav">
          <button class="btn btn_check"><iconCheck />Next</button>
        </nav>
      </dd>
    </dl>

    <!-- Step 02 -->
    <dl class="vow_make_item step02">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">【Step 02】誓いを入力してください。</span>
        </p>
      </dt>
      <dd>
        <div class="vow_make_input">
          <textarea name="vow_input" rows="4" cols="30" placeholder="誓いや目標をここに入力して、このeyevowに目入れしましょう"></textarea>
        </div>
        <nav class="vow_make_nav">
          <button class="btn btn_check"><iconCheck />Next</button>
        </nav>
      </dd>
    </dl>

    <!-- Step 03 -->
    <dl class="vow_make_item step03">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">【Step 03】この内容で誓いを立てますか？<br />※誓いは後から変更できないので、ちゃんと確認してね。</span>
        </p>
      </dt>
      <dd>
        <div class="vow_make_confirm">
          <p class="txt"><iconCheck />タイプ：イラスト</p>
          <p class="txt"><iconCheck />誓い：宇宙飛行士になるぞ！</p>
        </div>
        <nav class="vow_make_nav">
          <!-- <button class="btn btn_check"><iconCheck />Next</button> -->
          <router-link class="btn btn_vow" to="/eyevow"><iconEye />VOW</router-link>
        </nav>
      </dd>
    </dl>
  </div>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>


<script>
import 'swiper/dist/css/swiper.css'
import { swiper, swiperSlide } from 'vue-awesome-swiper'

import iconCheck from '@/assets/img/icon_check_line.svg?inline'
import iconEye from '@/assets/img/icon_eye.svg?inline'

export default {
  data() {
    return {
      swiperOption: {
        loop: true,
        centeredSlides: true,
        slidesPerView: 1.4,
        spaceBetween: 20,
        pagination: {
          el: '.swiper-pagination'
          // dynamicBullets: true
        }
      }
    }
  },
  components: {
    swiper,
    swiperSlide,
    iconCheck,
    iconEye
  }
}

</script>
