<template>
  <div class="about">
    <h1>This is an about page</h1>
  </div>
</template>

<script>

export default {
  created() {
    // display store data
    // console.log('about vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
