<template>
  <div class="product-detail">
    <h1>Product Detail</h1>
    <div v-if="item" key="product">
      <p>This page is ID. {{ id }} product!</p>
      <hr>
      <dl>
        <dt>Name: {{item.name}}</dt>
        <dd>Price: {{item.price}}</dd>
        <dd>Desc: {{item.desc}}</dd>
      </dl>
    </div>
  </div>
</template>

<script>
import products from '@/api/products.js'

export default {
  props: {
    'id': Number
  },
  data: function() {
    return {
      item: null
    }
  },
  watch: {
    id: {
      handler() {
        products.asyncFind(this.id, item => {
          this.item = item
        })
      }, immediate: true
    }
  }
}
</script>
