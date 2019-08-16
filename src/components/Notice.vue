<template>
  <div id="notice" class="notice">
    <nav class="notice_icon">
      <span class="badge"></span>
      <figure class="icon">
        <img alt="" src="@/assets/img/icon_notice.svg">
      </figure>
    </nav>
    <div class="notice_detail">
      <p class="notice_detail_item">
        <span class="icon icon_cheer"><img alt="" src="@/assets/img/icon_line_cheer.svg"></span>
        <span class="txt">あなたの誓いにCHEERが <span class="cheer_num">24</span> つつきました</span>
        <span class="time">2h</span>
      </p>
      <p class="notice_detail_item">
        <span class="icon icon_notice"><img alt="" src="@/assets/img/icon_line_notice.svg"></span>
        <span class="txt">eyevowの運営事務所から大切なお知らせがあります。eyevowの運営事務所から大切なお知らせがあります。</span>
        <span class="time">4h</span>
      </p>
      <p class="notice_detail_item">
        <span class="icon icon_cheer"><img alt="" src="@/assets/img/icon_line_cheer.svg"></span>
        <span class="txt">あなたの誓いにCHEERが <span class="cheer_num">3</span> つつきました</span>
        <span class="time">19d</span>
      </p>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Notice',
  components: {
  }
}
</script>

<style scoped>
  /* header.global {
    text-align: center;
    background: #ccc;
  } */
</style>
