<template>
  <div class="cheer">
    <h1>This is an cheer page</h1>
  </div>
</template>

<script>

export default {
  created() {
    // display store data
    console.log('cheer vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
