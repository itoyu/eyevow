<template>
  <div class="product-detail">
    <h1>Product Detail</h1>
    <p>This page is ID. {{ id }} product!</p>
  </div>
</template>

<script>
export default {
  props: {
    'id': Number
  }
}
</script>
