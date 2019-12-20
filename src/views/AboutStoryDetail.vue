<template>
  <div class="tab_contents_item">

    <ul class="tab_nav tab_nav_story">
<router-link to="/about/story/01" class="tab_nav_item" :class="{'active': (storyNum === '01')}">Ep.01</router-link>
<router-link to="/about/story/02" class="tab_nav_item" :class="{'active': (storyNum === '02')}">Ep.02</router-link>
<router-link to="/about/story/03" class="tab_nav_item" :class="{'active': (storyNum === '03')}">Ep.03</router-link>
<router-link to="/about/story/epilogue" class="tab_nav_item" :class="{'active': (storyNum === 'epilogue')}">Epilogue</router-link>
    </ul>

    <div class="story story_detail" v-if="storyNum === '01'">
      <p class="story_detail_tit"><span class="num">Ep.01</span>　アニミズム</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-1.jpg">
        </figure>
        <p class="txt">「モノ」には魂が宿ると言われている。<br>生き物から一定量の想いがモノに込められると、それは魂というカタチに昇華して、やがて新たな命になる。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-2.jpg">
        </figure>
        <!-- <p class="txt">しかし2080年、魂は失われかけていた…</p> -->
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-3.jpg">
        </figure>
        <p class="txt">技術と科学の発展により地球はとても豊かになった。<br>AIの指示とロボットの働きにより、資源と世界は効率的に運営されていた。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-4.jpg">
        </figure>
        <p class="txt">人類はDNA分析により、能力は全て数値化された。<br>そしてAIによるライフプランが提案され、誰もがそれに従い、満足した人生を送っていた。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-5.jpg">
        </figure>
        <p class="txt">こうして人類は争いもなく、平和な世界が訪れていた。<br><br>…と同時に人類の進化は終わりを迎えた。<br>変化を必要としなくなったため、人々は希望や目標を失い「魂」の存在は消えかけていた…</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-6.jpg">
        </figure>
        <!-- <p class="txt">しかし！</p> -->
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-7.jpg">
        </figure>
        <p class="txt">意思を持つ「モノ」達の世界である『Thing workd』では魂を必要としていた。<br>人々にもう一度希望を持たせるため、「モノ」達は立ち上がった！</p>
      </div>

      <p class="story_detail_tit"><span class="num">Ep.01</span>　完</p>
    </div>


    <div class="story story_detail" v-if="storyNum === '02'">
      <p class="story_detail_tit"><span class="num">Ep.02</span>　Thing World</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-1.jpg">
        </figure>
        <p class="txt">ここは意思を持つ「モノ」達の世界「Thing World」<br>モノ達は人々の希望や誓いを「魂の宿り木」に集め、原動力にしていた。<br><br>しかし、進化の止まった人類からの誓いは減少し、モノ達は困っていた。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-2.jpg">
        </figure>
        <p class="txt">このピンチに立ち上がったのが『縁起物族』だった。<br>古来よりたくさんの想いや願いを受け、魂を生み出したモノたちである。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-3.jpg">
        </figure>
        <p class="txt">縁起物の中でも、特に「だるま」の目に込められる"誓い"や"祈願"は最も生命力に溢れ、美しい魂を生むことができた。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-4.jpg">
        </figure>
        <p class="txt">そこで、村で一番の生命力を持つ「だるま」が地球に行って人類を奮起させることになったのであった。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-5.jpg">
        </figure>
        <p class="txt">だるまは、縁起物族の力で仮の姿に返信し、地球に降り立った。<br>「人類奮起計画」が今、幕を開けた。</p>
      </div>

      <p class="story_detail_tit"><span class="num">Ep.02</span>　完</p>
    </div>


    <div class="story story_detail" v-if="storyNum === '03'">
      <p class="story_detail_tit"><span class="num">Ep.03</span> 「eye, vow!!」</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-1.jpg">
        </figure>
        <p class="txt">無事に地球に降り立った「だるま」は早速、希望を抱いている人を探し始めた。<br>すると人組の親子に目が止まった。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-2.jpg">
        </figure>
        <p class="txt">
          子供：「お母さん、ぼく将来はサッカー選手になりたい！」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-3.jpg">
        </figure>
        <p class="txt">
          母：「…残念ながら、あなたの成長数値はプロの基準に達してないの。<br>だからサッカー選手は諦めて、研究職を目指しましょう。」<br><br>
          子供：「お母さん、ぼく将来はサッカー選手になりたい！」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-4.jpg">
        </figure>
        <p class="txt">
          だるま：「ちょっ、なぜそんなに早く諦めるんだ！？<br>挑戦もしてないのに諦める必要なんてない！」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-5.jpg">
        </figure>
        <p class="txt">だるま：「キミの希望、この目に込めてやるからちょっと待て…」</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-6.jpg">
        </figure>
        <p class="txt">『eye, vow!』</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-7.jpg">
        </figure>
        <p class="txt">
          母：「あれ、なんだか…サッカー選手に…なれる気がしてきたわ」<br>
          子供：「うん、僕もう少し頑張ってみるよ！ありがとう、おねーちゃん！」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-8.jpg">
        </figure>
        <p class="txt">だるま：「うん、いい誓いが生まれてよかった！<br>この調子で、他のみんなの誓いもワタシの目に込めていこう」<br><br>こうして、人々の誓いを瞳に宿し、魂を生み出す活動こそがThinb Worldを救う『eyevow』プロジェクトなのだ！</p>
      </div>

      <p class="story_detail_tit"><span class="num">Ep.03</span>　完</p>
    </div>



    <div class="story story_detail" v-if="storyNum === 'epilogue'">
      <p class="story_detail_tit"><span class="num">Epilogue</span></p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-1.jpg">
        </figure>
        <p class="txt">だるま：「あ、あれはなんと！誓い達成の光だ！」</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-2.jpg">
        </figure>
        <p class="txt">
          『Achieved!』
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-3.jpg">
        </figure>
        <p class="txt">
          だるま：「よかった…！これで私は魂に昇華できて、みんなも救われる」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-4.jpg">
        </figure>
        <p class="txt">
          招き猫：「よくやったな、おめでとう<br>これでお主にも新たな命が宿るだろう<br>さあ、宿り木の元へ行こう」
        </p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-5.jpg">
        </figure>
        <p class="txt">だるまが達成した誓いの想いを捧げると、だるまと宿り木は光に包まれた。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-6.jpg">
        </figure>
        <p class="txt">だるまはついに「魂」というカタチになった。<br>そして"希望の雨"をThing Worldに降らせながら空に昇っていった。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-7.jpg">
        </figure>
        <p class="txt">希望に満ちた、新たな命がまた1つ誕生した。</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/4-8.jpg">
        </figure>
        <p class="txt">だるま：「誓いをありがとう。キミとまた会えるのを楽しみにしているよ」</p>
      </div>

      <p class="story_detail_tit">Thank you!</p>
    </div>

    <ul class="tab_nav">
      <router-link to="/about/story" class="tab_nav_item">&lt; Back</router-link>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'AboutStoryDetail',
  data: function() {
    return {
      storyNum: ''
    }
  },
  methods: {
    storypath: function() {
      const pathSep = location.pathname.split('/story/');
      this.storyNum = pathSep[1];
    }
  },
  created() {
    this.storypath();
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.storypath();
  }
}
</script>
