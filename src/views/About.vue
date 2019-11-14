<template>
  <div class="about">
    <ul class="tab_nav">
      <router-link to="/about" class="tab_nav_item" exact>Tutorial</router-link>
      <router-link to="/about/story" class="tab_nav_item">Story</router-link>
      <router-link to="/about/statement" class="tab_nav_item">Statement</router-link>
      <router-link to="/about/credit" class="tab_nav_item">Credit</router-link>
    </ul>

    <div class="tab_contents">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>

export default {
  created() {
    console.log('about page!');
    this.init();
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.remove('focus');
    }
  }
}
</script>
