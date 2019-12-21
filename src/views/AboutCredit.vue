<template>

  <div class="tab_contents_item">
    <div class="credit">
      <h2 class="credit_tit">Member</h2>

      <ul class="credit_list">
        <li><a class="credit_list_item" href="https://kajil.co/" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile01.png"></figure>
          <div class="credit_info">
            <p class="credit_name">KAJIL</p>
            <p class="credit_role">企画 / 進行</p>
            <p class="credit_text">まずは英語と中国語に対応して世界中の人にeyevowを使ってもらえるようにしたい。あとはフィギュアとvtuberも生み出したい。</p>
          </div>
        </a></li>
        <li><a class="credit_list_item" href="http://tomoka.world" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile02.png"></figure>
          <div class="credit_info">
            <p class="credit_name">TOMOKA</p>
            <p class="credit_role">デフォルメイラスト / 撮影</p>
            <p class="credit_text">イラストのお仕事を今の二倍に！そしていつかアナザースカイでNYに！</p>
          </div>
        </a></li>
        <li><a class="credit_list_item" href="https://lunalia.tokyo" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile03.png"></figure>
          <div class="credit_info">
            <p class="credit_name">Lunà/ia</p>
            <p class="credit_role">衣装 / 製作全般</p>
            <p class="credit_text">本業の音楽家として生きる！ 希望に満ち溢れる、心が震える音楽をたくさんつくりたい。浴びたい。世界中に響かせたい。つまり音楽の仕事ください！</p>
          </div>
        </a></li>
        <li><a class="credit_list_item" href="https://twitter.com/to_co_pe" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile04.png"></figure>
          <div class="credit_info">
            <p class="credit_name">tocope</p>
            <p class="credit_role">キャラクターデザイン</p>
            <p class="credit_text">一年間無事</p>
          </div>
        </a></li>
        <li><a class="credit_list_item" href="https://suichu.net" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile06.png"></figure>
          <div class="credit_info">
            <p class="credit_name">合同会社水中工房</p>
            <p class="credit_role">サイト開発</p>
            <p class="credit_text">いのちだいじに</p>
          </div>
        </a></li>
        <li><a class="credit_list_item" href="https://instagram.com/arami5652" target="_blank">
          <figure class="credit_icon"><img alt="" src="@/assets/img/profile/profile05.png"></figure>
          <div class="credit_info">
            <p class="credit_name">荒木美帆</p>
            <p class="credit_role">レイヤー</p>
            <p class="credit_text">相手の気持ちにちゃんと寄り添える人(タレント)になる✨目標は・・嵐のみなさんと、頑張る誰かの背中を押してエールを送るバラエティ番組を一緒にする！</p>
          </div>
        </a></li>
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
