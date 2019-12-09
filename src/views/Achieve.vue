<template>
  <div class="vow">
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_02.png"></span>
        <span class="txt">なんと！あなた誓いを達成したんですか？</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_02.png"></span>
        <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：{{vow.type}}<br>■誓い：{{vow.text}}<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_02.png"></span>
        <span class="txt">完成の目入れをしましょう？</span>
      </p>
      <nav class="vow_make_nav center"> <!-- v-on:click="vowMakeHandlerCheck" -->
        <!-- <router-link class="btn btn_vow" to="/eyevow"></router-link> -->
        <button @click="showPopup" class="btn btn_vow"><iconEye />OK !</button>
      </nav>
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
              <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：{{vow.type}}<br>■誓い：{{vow.text}}<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
            </p>
          </dt>
          <dd>
            <nav class="vow_make_nav center"> <!-- v-on:click="vowMakeHandlerCheck" -->
              <button @click="hidePopup" class="btn btn_back">Back</button>
              <router-link  class="btn btn_vow" to="/achieved"><iconEye />Achieved !</router-link>
            </nav>
          </dd>
        </dl>
      </div>
    </div>
  </div>
</template>

<script>

import iconEye from '@/assets/img/icon_eye.svg?inline'

export default {
  name: 'Achieved',
  data() {
    return {
      vow: {
        type: 'illust',
        text: 'Text'
      }
    }
  },
  components: {
    iconEye
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
      document.querySelector('body').classList.remove('focus_vow');
    },
    // Confirm Popup
    showPopup: function() {
      document.querySelector('.popup').classList.add('show');
    },
    hidePopup: function() {
      document.querySelector('.popup').classList.remove('show');
    }
  },
  created() {
    this.init();
  }
}
</script>
