<template>
  <div class="tab_contents_item">

    <ul class="tab_nav tab_nav_story">
<router-link to="/about/story/01" class="tab_nav_item" :class="{'active': (storyNum === '01')}">Ep.01</router-link>
<router-link to="/about/story/02" class="tab_nav_item" :class="{'active': (storyNum === '02')}">Ep.02</router-link>
<router-link to="/about/story/03" class="tab_nav_item" :class="{'active': (storyNum === '03')}">Ep.03</router-link>
<router-link to="/about/story/epilogue" class="tab_nav_item" :class="{'active': (storyNum === 'epilogue')}">Epilogue</router-link>
    </ul>

    <div class="story story_detail" v-if="storyNum === '01'">
      <p class="story_detail_tit"><span class="num">Ep.01</span>「祈願」の消滅</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-1.jpg">
        </figure>
        <p class="txt">ストーリー: 1-1</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-2.jpg">
        </figure>
        <p class="txt">ストーリー: 1-2</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-3.jpg">
        </figure>
        <p class="txt">ストーリー: 1-3</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-4.jpg">
        </figure>
        <p class="txt">ストーリー: 1-4</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-5.jpg">
        </figure>
        <p class="txt">ストーリー: 1-5</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-6.jpg">
        </figure>
        <p class="txt">ストーリー: 1-6</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/1-7.jpg">
        </figure>
        <p class="txt">ストーリー: 1-7</p>
      </div>
    </div>

    <div class="story story_detail" v-if="storyNum === '02'">
      <p class="story_detail_tit"><span class="num">Ep.02</span>“VOWS WORLD”</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-1.jpg">
        </figure>
        <p class="txt">ストーリー: 2-1</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-2.jpg">
        </figure>
        <p class="txt">ストーリー: 2-2</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-3.jpg">
        </figure>
        <p class="txt">ストーリー: 2-3</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-4.jpg">
        </figure>
        <p class="txt">ストーリー: 2-4</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/2-5.jpg">
        </figure>
        <p class="txt">ストーリー: 2-5</p>
      </div>
    </div>

    <div class="story story_detail" v-if="storyNum === '03'">
      <p class="story_detail_tit"><span class="num">Ep.03</span>「eye, vow!!」</p>

      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-1.jpg">
        </figure>
        <p class="txt">ストーリー: 3-1</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-2.jpg">
        </figure>
        <p class="txt">ストーリー: 3-2</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-3.jpg">
        </figure>
        <p class="txt">ストーリー: 3-3</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-4.jpg">
        </figure>
        <p class="txt">ストーリー: 3-4</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-5.jpg">
        </figure>
        <p class="txt">ストーリー: 3-5</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-6.jpg">
        </figure>
        <p class="txt">ストーリー: 3-6</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-7.jpg">
        </figure>
        <p class="txt">ストーリー: 3-7</p>
      </div>
      <div class="story_detail_item">
        <figure class="img">
          <img alt="" src="@/assets/img/story/3-8.jpg">
        </figure>
        <p class="txt">ストーリー: 3-8</p>
      </div>
    </div>

    <div class="story story_detail" v-if="storyNum === 'epilogue'">
      <p class="story_detail_tit"><span class="num">Epilogue is </span>Coming soon.</p>
    </div>

    <ul class="tab_nav">
      <router-link to="/about/story" class="tab_nav_item">&lt; Story</router-link>
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
