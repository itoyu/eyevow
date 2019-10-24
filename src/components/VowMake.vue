<template>
  <div class="vow_make" id="vow_make">
    <!-- Step 01 -->
    <dl class="vow_make_item" id="vow_make_step01">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">まず、eyevowのタイプを選んで。</span>
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
          <toggle-button v-on:change="vowMakeHandler_type"
            :value="false" :labels="{checked: 'OK', unchecked: 'Pick'}" :width="64" :height="28" color="#3b469B" id="vowMake01" />
        </nav>
      </dd>
    </dl>

    <!-- Step 02 -->
    <dl class="vow_make_item disabled" id="vow_make_step02">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">最後にあなたの誓いを入力して。</span>
        </p>
      </dt>
      <dd>
        <div class="vow_make_input">
          <textarea v-model="vowText" :class="{'blank': (vowText === '')}"
           name="vow_input" rows="4" cols="30" placeholder="誓いや目標をここに入力して、このeyevowに目入れしましょう"></textarea>
        </div>
        <nav class="vow_make_nav">
          <toggle-button v-on:change="vowMakeHandler_text"
            :disabled="vowTextBlank" :value="false" :labels="{checked: 'OK', unchecked: 'Input'}" :width="64" :height="28" color="#3b469B" />
        </nav>
      </dd>
    </dl>

    <!-- Step 03 -->
    <dl class="vow_make_item disabled" id="vow_make_step03">
      <dt>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">これで準備完了。<br>↓の内容で誓いを立てていい？</span>
        </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：イラスト<br>■誓い：宇宙飛行士になるぞぉ！<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
          </p>
        <p class="chat_item">
          <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
          <span class="txt">ちなみに、誓いは後から変更できないから。<br>上に戻って直せるのは今だけ。</span>
        </p>
      </dt>
      <dd>
        <nav class="vow_make_nav center">
          <router-link class="btn btn_vow" to="/eyevow"><iconEye />OK !</router-link>
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

import iconEye from '@/assets/img/icon_eye.svg?inline'

export default {
  data() {
    return {
      vowTextData: '',
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
  methods: {
    init: function() {
      // const vowMakeItems = document.querySelectorAll('.vow_make_item');
      const vowMakeType = document.getElementById('vow_make_step01');
            // vowMakeText = document.getElementById('vow_make_step02'),
            // vowMakeConfirm = document.getElementById('vow_make_step03');

      // console.log(vowMakeType);
      // console.log(vowMakeText);
      // console.log(vowMakeConfirm);


      // setTimeout(function() {
      //   console.log('detachEvents');
      //   swiper.classList.add('disabled');
      // }, 5000);


      function vowMakeStart() {
        vowMakeType.classList.add('show');
      }

      // Start
      setTimeout(function() {
        vowMakeStart();
      }, 1000);
    },

    // Step01
    vowMakeHandler_type: function(e) {
      const vowMakeType = document.getElementById('vow_make_step01'),
            vowMakeText = document.getElementById('vow_make_step02');

      (e.value) ? disabledPick() : enablePick();

      function disabledPick() {
        vowMakeType.classList.add('disabled');

        // Enable Step02
        setTimeout(function() {
          vowMakeText.classList.add('show');
          vowMakeText.classList.remove('disabled');
        }, 500);
      }
      function enablePick() {
        vowMakeType.classList.remove('disabled');

        // Disable Step02
        setTimeout(function() {
          vowMakeText.classList.add('disabled');
        }, 500);
      }
    },

    // Step02
    vowMakeHandler_text: function(e) {
      const vowMakeText = document.getElementById('vow_make_step02'),
            vowMakeConfirm = document.getElementById('vow_make_step03');

      (e.value) ? disabledPick() : enablePick();

      function disabledPick() {
        vowMakeText.classList.add('disabled');

        // Enable Step03
        setTimeout(function() {
          vowMakeConfirm.classList.add('show');
          vowMakeConfirm.classList.remove('disabled');
        }, 500);
      }
      function enablePick() {
        vowMakeText.classList.remove('disabled');

        // Disable Step03
        setTimeout(function() {
          vowMakeConfirm.classList.add('disabled');
        }, 500);
      }

      console.log(this.vowText);
    }
  },
  computed: {
    vowText: {
      get() {
        return this.vowTextData
      },
      set(value) {
        this.vowTextData = value
      }
    },
    vowTextBlank() {
      return (this.vowTextData === '') ? true : false;
    }
  },
  components: {
    swiper,
    swiperSlide,
    iconEye
  },
  mounted() {
    this.init();
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
</style>
