<template>
  <div class="vow">
    <h1 class="page_title">Vow</h1>
    <Transform />
    <VowComplete />
  </div>
</template>

<script>
// import Transform from '@/components/Transform.vue'
import VowMake from '@/components/VowMake.vue'

export default {
  name: 'Vow',
  components: {
    Transform,
    VowComplete
  }
}
</script>
