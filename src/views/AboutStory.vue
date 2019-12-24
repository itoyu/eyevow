<template>

  <div class="tab_contents_item">
    <div class="story">
      <router-link to="/about/story/01" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/story/1-3.jpg">
        </figure>
        <div class="story_info">
          <p class="label"><span class="num">Ep.01</span>アニミズム</p>
          <p class="txt">2080年、AIとロボットの発展により人類の進化は終わりを迎えた。</p>
        </div>
      </router-link>
      <router-link to="/about/story/02" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/story/2-1.jpg">
        </figure>
        <div class="story_info">
          <p class="label"><span class="num">Ep.02</span>Thing World</p>
          <p class="txt">ここは生き物からの希望や誓いを原動力にしている「モノ」たちの世界「Thing World」</p>
        </div>
      </router-link>
      <router-link to="/about/story/03" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/story/3-5.jpg">
        </figure>
        <div class="story_info">
          <p class="label"><span class="num">Ep.03</span>eye, vow !!</p>
          <p class="txt">星からの司令で、人類に再び「誓い」を与えるために地球に降り立ったヒロイン</p>
        </div>
      </router-link>
      <!--
      <router-link to="/about/story/epilogue" class="story_item">
        <figure class="story_img">
          <img alt="" src="https://dummyimage.com/1000x1000/333333/ffffff&text= SECRET">
        </figure>
        <div class="story_info">
          <p class="label">Epilogue</p>
          <p class="txt note">※このエピソードは誓い達成後に見ることができます。</p>
        </div>
      </router-link>
      -->
    </div>
  </div>

</template>

<script>

export default {
  name: 'AboutStory',
  components: {
  },
  created() {
    this.init();
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
      document.querySelector('body').classList.remove('focus_vow');
    }
  }
}
</script>
