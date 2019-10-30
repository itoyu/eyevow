<template>
  <div class="transform">
    <div class="transform_in">
      <figure class="transform_anim">
        <p id="trans_anim_gif"><img alt="" src="@/assets/img/transform01.gif"></p>
        <p>Transform Animation!</p>
      </figure>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {

    }
  },
  components: {
  },
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
          console.log(result);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    startTransform: function() {
      document.querySelector('.transform').classList.add('show');

      // done!
      window.setTimeout(() => {
        this.endTransform();
      }, 4000);
    },
    endTransform: function() {
      document.querySelector('.transform').classList.remove('show');
      document.querySelector('.transform').classList.add('hide');
    }
  },
  mounted() {
    this.startTransform();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
</style>
