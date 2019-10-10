<template>
  <div class="vow">
    <h1 class="page_title">Vow</h1>
    <Transform />
    <VowComplete />
  </div>
</template>

<script>
import Transform from '@/components/Transform.vue'
import VowComplete from '@/components/VowComplete.vue'


export default {
  name: 'Eyevow',
  components: {
    Transform,
    VowComplete
  }
}
</script>
