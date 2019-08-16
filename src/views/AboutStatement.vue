<template>
  <div class="about">
    <h1 class="page_title">About</h1>

    <ul class="tab_nav">
      <router-link to="/about" class="tab_nav_item">Tutorial</router-link>
      <router-link to="/about/story" class="tab_nav_item">Story</router-link>
      <router-link to="/about/statement" class="tab_nav_item active">Statement</router-link>
      <router-link to="/about/credit" class="tab_nav_item">Credit</router-link>
    </ul>

    <div class="tab_contents">
      <div class="tab_contents_item">

        <div class="statement">
          <h2 class="statement_tit">Statement</h2>
          <h2 class="statement_catch">自分を切り開け</h2>
          <div class="statement_txt">
            <p>このプロジェクト誓いによって”夢”と”覚悟”を問いかけるプロジェクトです。</p>
            <p>近い将来、遺伝子の研究と解明が進んで人の限界値が見えてしまうと思います。</p>
            <p>AIが人間の知能を超え、ロボットが人間の肉体を超えた時、人が自ら何かをすることに意味を見出せるでしょうか？</p>
            <p>私たちは絶対にあると信じています。<br>一人一人が誓いを立てて、<br>夢を叶える世の中にしましょう。</p>
          </div>
        </div>

      </div>
    </div>

    <StartBtn />
    
  </div>
</template>

<script>
import StartBtn from '@/components/StartBtn.vue'

export default {
  name: 'AboutTutorial',
  components: {
    StartBtn
  }
}
</script>
