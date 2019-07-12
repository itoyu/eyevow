<template>
  <div class="product-list">
    <h1>Product List</h1>
  </div>
</template>

<script>

export default {
  created() {
    // display store data
    console.log(this.$store.state.count);
  }
}
</script>
