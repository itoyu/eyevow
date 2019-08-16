<template>
  <div class="about">
    <h1 class="page_title">About</h1>

    <ul class="tab_nav">
      <router-link to="/about" class="tab_nav_item">Tutorial</router-link>
      <router-link to="/about/story" class="tab_nav_item">Story</router-link>
      <router-link to="/about/statement" class="tab_nav_item">Statement</router-link>
      <router-link to="/about/credit" class="tab_nav_item active">Credit</router-link>
    </ul>

    <div class="tab_contents">
      <div class="tab_contents_item">

        <div class="credit">
          <h2 class="credit_tit">Member</h2>
          <dl class="credit_item">
            <dt>Character Design</dt>
            <dd><a href="#" target="_blank">tocope</a></dd>
            <dt>Illustration</dt>
            <dd><a href="#" target="_blank">tocope</a></dd>
            <dt>Cosplay</dt>
            <dd><a href="#" target="_blank">XXXX</a></dd>
            <dt>Photo</dt>
            <dd><a href="#" target="_blank">XXXX</a></dd>
            <!-- <dt>Site Design</dt>
            <dd><a href="#" target="_blank">tocope</a></dd> -->
            <dt>Develop</dt>
            <dd><a href="#" target="_blank">suichu.inc</a></dd>
            <dt>Coordinate</dt>
            <dd><a href="#" target="_blank">TOMOKA, Lunà/ia</a></dd>
            <dt>Logo</dt>
            <dd><a href="#" target="_blank">Shin rei</a></dd>
            <dt>Translation</dt>
            <dd><a href="#" target="_blank">Kimi tang</a></dd>
            <dt>Produce</dt>
            <dd><a href="#" target="_blank">KAJIL</a></dd>
          </dl>

          <h2 class="credit_tit">Resource</h2>
          <dl class="credit_item">
            <dt>Icon</dt>
            <dd><a href="https://www.flaticon.com" target="_blank">Freepik</a></dd>
          </dl>

        </div>

      </div>
    </div>

  </div>
</template>

<script>

export default {
  name: 'AboutCredit',
  created() {
    // display store data
    // console.log('about vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
