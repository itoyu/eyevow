<template>
  <div class="vow">
    <VowComplete />
    <Transform />
  </div>
</template>

<script>
// import PopConfirm from '@/components/PopConfirm.vue'
import VowComplete from '@/components/VowComplete.vue'
import Transform from '@/components/Transform.vue'


export default {
  name: 'Eyevow',
  components: {
    // PopConfirm,
    VowComplete,
    Transform
  },
  methods: {
    init: function() {
      document.querySelector('body').classList.add('focus');
    }
  },
  created() {
    this.init();
  }
}
</script>
