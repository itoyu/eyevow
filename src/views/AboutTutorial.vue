<template>
  <div class="tab_contents_item">

    <div class="tutorial">

      <div class="tutorial_box">
        <h2 class="statement_tit">What's eyevow</h2>

        <div class="page_lead">
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">
              日本では、古くから誓いを立てる時に縁起物である「だるま」に目入れをする文化があるらしい。<br>
              <span class="img"><img alt="" src="@/assets/img/about/about01.png"></span>
            </span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">そんな「だるま」が、人々の誓いを応援するために人の姿になったのが私「eyevow」だよ。</span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">私は、あなたの誓いを瞳に込めることで、“相棒”となり、あなたの誓いを応援するよ。<br>
              <span class="img"><img alt="" src="@/assets/img/about/about02.png"></span>
            </span>
          </p>
        </div>
      </div>


      <div class="tutorial_box">
        <h2 class="statement_tit">How to</h2>

        <div class="page_lead">
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt"><span class="step">Step.01</span><br>
              まず、目入れをするeyevowのタイプを選んで。<br>イラストか写真が選べるよ。<br>
              <span class="img"><img alt="" src="@/assets/img/about/about03.png"></span>
            </span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt"><span class="step">Step.02</span><br>
               次に、あなたの誓いを入力して。<br>
              <span class="img"><img alt="" src="@/assets/img/about/about04.png"></span>
            </span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt"><span class="step">Step.03</span><br>
              そうすると1回目の「目入れ」完了だよ。<br>
              <span class="img"><img alt="" src="@/assets/img/about/about05.png"></span>
            </span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt">
              あとは達成に向けてがんばってね。
            </span>
          </p>
          <p class="chat_item">
            <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
            <span class="txt"><span class="step">Step.04</span><br>
              そしていよいよ誓いを達成したら教えてね。<br>何が起こるかは、達成してからのお楽しみ。がんばってねー。
            </span>
          </p>
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
    // this.$store.commit('countup')
  }
}
</script>
