<template>
  <div class="transform">
    <div class="transform_in">
      <p>transform_in</p>
    </div>
    <div class="transform_achieved">
      <p>transform_achieved</p>
    </div>
  </div>
</template>

<script>

export default {
  methods: {
    init: function() {

      const promise = new Promise((resolve) => {
        window.setTimeout(() => {
          console.log('step01');
        }, 1000);

        window.setTimeout(() => {
          resolve('step02');
        }, 5000);

      });

      promise
        .then((result) => {
          console.log('result');
          console.log(result);
        })
        .catch((error) => {
          console.log(error);
        });
    }
  },
  created() {
    this.init();
  }
}
</script>
