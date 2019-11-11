<template>
  <div class="vow">
    <AchieveComplete />
    <Transform />
  </div>
</template>

<script>
// import PopConfirm from '@/components/PopConfirm.vue'
import AchieveComplete from '@/components/AchieveComplete.vue'
import Transform from '@/components/Transform2.vue'


export default {
  name: 'Eyevow',
  components: {
    // PopConfirm,
    AchieveComplete,
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
