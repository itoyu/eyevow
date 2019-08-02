<template>
  <div class="edit">
    <h1>This is an edit page</h1>
  </div>
</template>

<script>

export default {
  created() {
    // display store data
    // console.log('edit vue:' + this.$store.state.count);

    // count up store counter(commit method)
    this.$store.commit('countup')
  }
}
</script>
