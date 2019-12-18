<template>
  <div class="tab_contents_item">

    <div class="statement">
      <h2 class="statement_tit">Statement</h2>
      <h2 class="statement_catch">
        Humans, <br>be ambitious
        <!-- copyStatement / -->
      </h2>
      <div class="statement_txt">
        <p>これは”誓い”を通して<br>人に夢と覚悟を問いかけるプロジェクトです。</p>
        <p>AIやロボットの台頭によって人類は<br>存在意義を見つめ直されています。</p>
        <p>人類がデータや効率に囚われずに夢を見出すために、1人1人が多様な未来を描くために、あなたの背中を少しでも押すことができればこのプロジェクトは成功だと考えています。</p>
      </div>
    </div>
  </div>
</template>

<script>

import copyStatement from '@/assets/img/about/copy.svg?inline'

export default {
  name: 'AboutTutorial',
  components: {
    copyStatement
  }
}
</script>
