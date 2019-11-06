<template>
  <div class="vow">
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">じゃあいよいよ誓いを立てる</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">これから私の片目に、あなたの誓いを込める</span>
      </p>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Achieved',
  components: {
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
    }
  },
  created() {
    this.init();
  }
}
</script>
