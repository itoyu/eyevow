<template>
  <div class="transform">
    <div class="transform_in">
      <div class="transform_anim">
        <div id="trans_anim_gif" class="transform_parts transform_bg"><img alt="" src="@/assets/img/effect_trans_bg.gif"></div>
        <figure class="transform_char01 transform_char transform_parts"><img alt="" src="@/assets/img/eyevow/character_cosplayers_01.png"></figure>
        <span class="transform_effect01 transform_parts"><img alt="" src="@/assets/img/effect_trans0101.gif"></span>
        <span class="transform_effect02 transform_parts"><img alt="" src="@/assets/img/effect_trans0102.gif"></span>
        <!-- <span class="transform_effect03 transform_parts"><img alt="" src="@/assets/img/effect_trans0103.gif"></span>
        <figure class="transform_char02 transform_char transform_parts"><img alt="" src="@/assets/img/eyevow/character_cosplayers_02.png"></figure> -->
      </div>
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
      }, 4500);

      // プロトタイプ表示
      window.setTimeout(() => {
        document.querySelector('.transform_char01').classList.add('ready');
      }, 1000);

      window.setTimeout(() => {
        document.querySelector('.transform_effect01').classList.add('ready');
        document.querySelector('.transform_effect02').classList.add('ready');
      }, 2000);


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
@keyframes charactor01 {
  0% {
    transform: translate(-50%, -10px);
  }
  70% {
    transform: translate(-50%, 0px);
  }
  100% {
    transform: translate(-50%, -10px);
  }
}

.transform {
  transform-style:preserve-3d;

  &_anim {

  }
  &_parts {
    position: absolute;
    transition: .5s;
  }
  &_char {
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    width: 100%;
  }
  .transform_bg {
    top: 50%;
    transform: translateY(-50%);
  }
  .transform_char01 {
    opacity: 0;
    transform: translate(-40%, -60%) rotate(10deg);
    filter: blur(1.5rem);


    &.ready {
      opacity: 1;
      filter: blur(0);
      transform: translate(-50%, -50%) rotate(0deg);
    }
  }
  .transform_char02 {
    opacity: 0;
  }
  .transform_effect01,
  .transform_effect02 {
    opacity: 0;

    &.ready {
      opacity: 1;
    }
  }
  .transform_effect01 {
    top: 15%;
    left: -5%;
    width: 200px;
  }
  .transform_effect02 {
    bottom: 15%;
    right: 5%;
    width: 200px;
  }
  .transform_effect03 {
    opacity: 0;

    &.ready {
      display: block;
      opacity: 1;
      left: 0;
      top: 0;
      width: auto;
      height: 100vh;
    }
  }
}
</style>
