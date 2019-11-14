<template>

  <div class="tab_contents_item">
    <div class="credit">
      <h2 class="credit_tit">Member</h2>

      <ul class="credit_list">
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">KAJIL</p>
            <p class="credit_role">企画 / 進行</p>
            <p class="credit_text">多くの人にeyevowを使ってもらいたい。</p>
          </div>
        </li>
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">TOMOKA</p>
            <p class="credit_role">XX / XX</p>
            <p class="credit_text">誓いをいれてもらう</p>
          </div>
        </li>
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">Lunà/ia</p>
            <p class="credit_role">XX / XX</p>
            <p class="credit_text">誓いをいれてもらう</p>
          </div>
        </li>
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">tocope</p>
            <p class="credit_role">XX / XX</p>
            <p class="credit_text">誓いをいれてもらう</p>
          </div>
        </li>
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">Arakawa</p>
            <p class="credit_role">XX / XX</p>
            <p class="credit_text">誓いをいれてもらう</p>
          </div>
        </li>
        <li class="credit_list_item">
          <figure class="credit_icon"><img alt="" src="@/assets/img/eyevow/icon_illust_01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">Toyama</p>
            <p class="credit_role">XX / XX</p>
            <p class="credit_text">誓いをいれてもらう</p>
          </div>
        </li>
      </ul>

      <h2 class="credit_tit">Resource</h2>
      <dl class="credit_item">
        <dt>Icon</dt>
        <dd><a href="//www.flaticon.com" target="_blank">Freepik</a></dd>
        <dt>Animation, Movie</dt>
        <dd><a href="//www.motionelements.com/" target="_blank">MotionElements</a></dd>
      </dl>

    </div>

  </div>
</template>

<script>

export default {
  name: 'AboutCredit',
  created() {
    // display store data
    // console.log('about vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
