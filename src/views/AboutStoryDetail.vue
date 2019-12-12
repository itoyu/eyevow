<template>
  <div class="tab_contents_item">

    <ul class="tab_nav tab_nav_story">
      <router-link to="/about/story" class="tab_nav_item">&lt; Story</router-link>
      <router-link to="/about/story/01" class="tab_nav_item active">Ep.01</router-link>
      <router-link to="/about/story/02" class="tab_nav_item">Ep.02</router-link>
      <router-link to="/about/story/03" class="tab_nav_item">Ep.03</router-link>
      <router-link to="/about/story/epilogue" class="tab_nav_item">Epilogue</router-link>
    </ul>

    <div class="story story_detail">
      <p class="story_detail_tit"><span class="num">Ep.01</span>「祈願」の消滅</p>

      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image01">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image02">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image03">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image04">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image05">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image06">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image07">
      </figure>
      <figure class="story_detail_img">
        <img alt="" src="https://dummyimage.com/700x400/ffffff/cccccc&text=Story+Image08">
      </figure>


      <ul class="tab_nav">
        <router-link to="/about/story" class="tab_nav_item">&lt; Story</router-link>
        <router-link to="/about/story/02" class="tab_nav_item">Ep.02 &gt;</router-link>
      </ul>
    </div>

  </div>
</template>

<script>

export default {
  name: 'AboutStoryDetail',
  computed: {
    storypath: function() {
      console.log(location.pathname);
      var pathSep = location.pathname.split('/story/');
      var storyNum = pathSep[1];

      console.log(pathSep);
      console.log(storyNum);

      return storyNum;
    }
  },
  created() {
    this.storypath();
  }
}
</script>
