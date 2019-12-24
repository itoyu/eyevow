<template>
  <div id="notice" class="notice">
    <button class="notice_icon" @click="tglNotice">
      <!-- <span class="badge"></span> -->
      <figure class="icon">
        <img alt="" src="@/assets/img/common/icon_notice.svg">
      </figure>
    </button>
    <div class="notice_detail" v-show="noticeShow">
      <p class="notice_detail_item">
        <span class="icon icon_cheer"><img alt="" src="@/assets/img/common/icon_line_cheer.svg"></span>
        <span class="txt">誓いを立てたので、一部のアイテムが<a href="https://eyevow.stores.jp/" class="txtlink" target="_blank">ECサイト</a>から買えるようになりました。</span>
        <span class="time">2h</span>
      </p>
      <p class="notice_detail_item">
        <span class="icon icon_notice"><img alt="" src="@/assets/img/common/icon_line_notice.svg"></span>
        <span class="txt">eyevowの運営事務所から大切なお知らせがあります。eyevowの運営事務所から大切なお知らせがあります。</span>
        <span class="time">4h</span>
      </p>
      <p class="notice_detail_item">
        <span class="icon icon_cheer"><img alt="" src="@/assets/img/common/icon_line_cheer.svg"></span>
        <span class="txt">あなたの誓いにCHEERが <span class="cheer_num">3</span> つつきました</span>
        <span class="time">19d</span>
      </p>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Notice',
  data() {
    return {
      noticeShow: false
    }
  },
  components: {
  },
  methods: {
    tglNotice: function() {
      this.noticeShow = !this.noticeShow;
    }
  }
}
</script>

<style scoped>
  /* header.global {
    text-align: center;
    background: #ccc;
  } */
</style>
