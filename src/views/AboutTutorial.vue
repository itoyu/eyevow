<template>
  <div class="tab_contents_item">

    <div class="tutorial">
      <div class="tutorial_item txt">
        <p>あなたが誓いを立て、あなたが実現をする、あなたの物語を問いかけ・見守るサービスです。<br>「eyevow」というキャラクターが、あなたの誓い達成を、目入れという行為で表現して応援します。</p>
      </div>

      <div class="tutorial_item">
        <figure class="tutorial_img">
          <img alt="" src="@/assets/img/about/img_tutorial01.png">
        </figure>
        <div class="tutorial_info">
          <p class="label">Step.01</p>
          <p class="txt">[イラスト] or [レイヤー] のどちらかのeyevowタイプを選びます。</p>
        </div>
      </div>
      <div class="tutorial_item">
        <figure class="tutorial_img">
          <img alt="" src="@/assets/img/about/img_tutorial02.png">
        </figure>
        <div class="tutorial_info">
          <p class="label">Step.02</p>
          <p class="txt">あなたの達成したい「誓い」を入力します。</p>
        </div>
      </div>
      <div class="tutorial_item">
        <figure class="tutorial_img">
          <img alt="" src="@/assets/img/about/img_tutorial03.png">
        </figure>
        <div class="tutorial_info">
          <p class="label">Step.03</p>
          <p class="txt">入力した「誓い」がeyevowの片目に込められ、目入れ完了です。</p>
        </div>
      </div>

      <div class="tutorial_item txt">
        <p>誓い達成を目指して、日々活動しましょう。</p>
      </div>

      <div class="tutorial_item">
        <figure class="tutorial_img">
          <img alt="" src="@/assets/img/about/img_tutorial04.png">
        </figure>
        <div class="tutorial_info">
          <p class="label">Step.04</p>
          <p class="txt">誓いを達成したら「Achived」ボタンを押します。</p>
        </div>
      </div>
      <div class="tutorial_item">
        <figure class="tutorial_img">
          <img alt="" src="@/assets/img/about/img_tutorial05.png">
        </figure>
        <div class="tutorial_info">
          <p class="label">Step.05</p>
          <p class="txt">もう片方に目入れを行い、eyevowが完全体となります。</p>
        </div>
      </div>
    </div>

  </div>
</template>

<script>

export default {
  name: 'AboutTutorial',
  components: {
  },
  created() {
    // display store data
    // console.log('about vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
