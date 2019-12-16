<template>

  <div class="tab_contents_item">
    <div class="story">
      <router-link to="/about/story/01" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/story/1-3.jpg">
        </figure>
        <div class="story_info">
          <p class="label"><span class="num">Ep.01</span>「祈願」の消滅</p>
          <p class="txt">21xx年、技術と科学の発展により人類の進化は遂に終わりを迎えた。</p>
        </div>
      </router-link>
      <router-link to="/about/story/02" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/story/2-1.jpg">
        </figure>
        <div class="story_info">
          <p class="label"><span class="num">Ep.02</span>“VOWS WORLD”</p>
          <p class="txt">ここは宇宙中の「希望」と「誓い」を原動力にしている星、“VOWS WORLD” </p>
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
      <router-link to="/about/story/epilogue" class="story_item">
        <figure class="story_img">
          <img alt="" src="@/assets/img/about/img_storyEp.png">
        </figure>
        <div class="story_info">
          <p class="label">Epilogue</p>
          <p class="txt note">※このエピソードは誓い達成後に見ることができます。</p>
        </div>
      </router-link>
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
