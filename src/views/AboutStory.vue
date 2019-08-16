<template>
  <div class="about">
    <h1 class="page_title">About</h1>

    <ul class="tab_nav">
      <router-link to="/about" class="tab_nav_item">Tutorial</router-link>
      <router-link to="/about/story" class="tab_nav_item active">Story</router-link>
      <router-link to="/about/statement" class="tab_nav_item">Statement</router-link>
      <router-link to="/about/credit" class="tab_nav_item">Credit</router-link>
    </ul>

    <div class="tab_contents">
      <div class="tab_contents_item">
        <!-- <h2 class="tab_tit">Story</h2> -->

        <div class="story">

          <router-link to="/about/story/01" class="story_item">
            <figure class="story_img">
              <img alt="" src="@/assets/img/about/img_story01.png">
            </figure>
            <div class="story_info">
              <p class="label"><span class="num">Ep.01</span>「祈願」の消滅</p>
              <p class="txt">21xx年、技術と科学の発展により人類の進化は遂に終わりを迎えた。</p>
            </div>
          </router-link>
          <router-link to="/about/story/02" class="story_item">
            <figure class="story_img">
              <img alt="" src="@/assets/img/about/img_story02.png">
            </figure>
            <div class="story_info">
              <p class="label"><span class="num">Ep.02</span>“VOWS WORLD”</p>
              <p class="txt">ここは宇宙中の「希望」と「誓い」を原動力にしている星、“VOWS WORLD” </p>
            </div>
          </router-link>
          <router-link to="/about/story/03" class="story_item">
            <figure class="story_img">
              <img alt="" src="@/assets/img/about/img_story03.png">
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
    </div>

    <StartBtn />

  </div>
</template>

<script>
import StartBtn from '@/components/StartBtn.vue'

export default {
  name: 'AboutStory',
  components: {
    StartBtn
  }
}
</script>
