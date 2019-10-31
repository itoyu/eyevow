<template>
  <div>
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
              <swiper-slide data-vowtype="illust">
                <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_01.png"></figure>
                <p class="txt">イラスト</p>
              </swiper-slide>
              <swiper-slide data-vowtype="photo">
                <figure class="img"><img alt="" src="@/assets/img/eyevow/character_photo_01.png"></figure>
                <p class="txt">写真</p>
              </swiper-slide>
            </swiper>
            <div class="swiper-pagination"></div>
          </div>
          <nav class="vow_make_nav">
            <toggle-button v-on:change="vowMakeHandler_type"
              :value="false" :labels="{checked: 'OK', unchecked: 'Pick'}" :width="64" :height="28" color="#3b469B" />
          </nav>
        </dd>
      </dl>

      <!-- Step 02 -->
      <dl class="vow_make_item" id="vow_make_step02">
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
      <dl class="vow_make_item" id="vow_make_step03">
        <dt>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">これで準備完了。<br>↓の内容で誓いを立てていい？</span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：{{vowMake.type}}<br>■誓い：{{vowMake.text}}<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">ちなみに、誓いは後から変更できないから。<br>上に戻って直せるのは今だけ。</span>
          </p>
        </dt>
        <dd>
          <nav class="vow_make_nav center"> <!-- v-on:click="vowMakeHandlerCheck" -->
            <!-- <router-link class="btn btn_vow" to="/eyevow"></router-link> -->
            <button @click="showPopup" :class="{'disabled': vowMakeReady}" class="btn btn_vow"><iconEye />OK !</button>
          </nav>
        </dd>
      </dl>
    </div>
    <div class="popup popup_confirm">
      <div class="popup_in">
        <dl>
          <dt>
            <p class="chat_item">
              <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
              <span class="txt">じゃあいよいよ目入れするわ。</span>
            </p>
            <p class="chat_item">
              <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
              <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：{{vowMake.type}}<br>■誓い：{{vowMake.text}}<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
            </p>
          </dt>
          <dd>
            <nav class="vow_make_nav center"> <!-- v-on:click="vowMakeHandlerCheck" -->
              <button @click="hidePopup" class="btn btn_back">Back</button>
              <router-link  class="btn btn_vow" to="/eyevow"><iconEye />Vow !</router-link>
            </nav>
          </dd>
        </dl>
      </div>
    </div>
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
      vowMake: {
        step01: false,
        step02: false,
        type: '',
        text: ''
      },
      swiperOption: {
        // loop: true,
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
      const vowMakeType = document.getElementById('vow_make_step01');

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
            vowMakeText = document.getElementById('vow_make_step02'),
            sliderPickItem = document.querySelector('.swiper-slide-active');

      if(e.value) {
        vowMakeType.classList.add('disabled');
        this.vowMake.type = sliderPickItem.getAttribute('data-vowtype');
        this.vowMake.step01 = true;

        // Enable Step02
        vowMakeText.classList.add('show');
      } else {
        vowMakeType.classList.remove('disabled');
        this.vowMake.step01 = false;
      }
    },

    // Step02
    vowMakeHandler_text: function(e) {
      const vowMakeText = document.getElementById('vow_make_step02'),
            vowMakeConfirm = document.getElementById('vow_make_step03');

      if(e.value) {
        vowMakeText.classList.add('disabled');

        this.vowMake.text = this.vowText;
        this.vowMake.step02 = true;

        // Enable Step03
        vowMakeConfirm.classList.add('show');
      } else {
        vowMakeText.classList.remove('disabled');
        this.vowMake.step02 = false;
      }
    },

    // Check
    vowMakeHandlerCheck: function() {
      console.log(this.vowMake.type);
      console.log(this.vowMake.text);
    },

    // Confirm Popup
    showPopup: function() {
      document.querySelector('.popup').classList.add('show');
    },
    hidePopup: function() {
      document.querySelector('.popup').classList.remove('show');
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
    },
    vowMakeReady() {
      return (this.vowMake.step01 && this.vowMake.step02) ? false : true;
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
