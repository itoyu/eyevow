<template>
  <div class="product-list">
    <h1>Product List</h1>
    <ul>
      <li v-for="item in productList" v-bind:key="item.id">
        <router-link :to="`/product/${ item.id }`">{{ item.name }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
import products from '@/api/products.js'

export default {
  computed: {
    productList: () => products.fetch()
  }
}
</script>
