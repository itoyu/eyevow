<template>
  <div class="vow">
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">じゃあいよいよ誓いを立ててみよう。</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">私の片目に、あなたの誓いを込めるわ。</span>
      </p>
    </div>
    <VowMake />
  </div>
</template>

<script>
import VowMake from '@/components/VowMake.vue'

export default {
  name: 'Vow',
  components: {
    VowMake
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
      document.querySelector('body').classList.add('focus_vow');
    }
  },
  created() {
    this.init();
  }
}
</script>
