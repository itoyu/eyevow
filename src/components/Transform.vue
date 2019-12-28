<template>
  <div class="transform">
    <div class="transform_in">
      <div class="transform_anim">
        <div id="trans_anim_gif" class="transform_parts transform_bg"><img alt="" src="@/assets/img/effect_trans_bg.gif"></div>
        <figure class="transform_char transform_parts">
          <span class="img img01"><img alt="" src="@/assets/img/story/3-5.jpg"></span>
          <span class="img img02"><img alt="" src="@/assets/img/story/3-6.jpg"></span>
          <span class="img img03 trans_eyevow" id="trans_eyevow01" :class="myVow.type"></span>
          <span class="img img04 trans_eyevow" id="trans_eyevow02" :class="myVow.type"></span>
        </figure>
        <span class="transform_effect01 transform_parts"><img alt="" src="@/assets/img/effect_card.gif"></span>
        <span class="transform_effect02 transform_parts"><img alt="" src="@/assets/img/effect_card.gif"></span>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  data: function() {
    return {
      myVow: {
        type: this.$store.getters.getMyVow.type,
        text: this.$store.getters.getMyVow.text
      }
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
      }, 6000);

      // プロトタイプ表示
      window.setTimeout(() => {
        document.querySelector('.transform_char').classList.add('step01');
      }, 1000);
      window.setTimeout(() => {
        document.querySelector('.transform_char').classList.add('step02');
      }, 2000);
      window.setTimeout(() => {
        document.querySelector('.transform_char').classList.add('step03');
      }, 3000);
      window.setTimeout(() => {
        document.querySelector('.transform_char').classList.add('step04');
      }, 4000);

      window.setTimeout(() => {
        document.querySelector('.transform_effect01').classList.add('ready');
        document.querySelector('.transform_effect02').classList.add('ready');
      }, 500);
    },
    endTransform: function() {
      // document.querySelector('.transform').classList.remove('show');
      document.querySelector('.transform').classList.add('white');
      window.setTimeout(() => {
        document.querySelector('.transform').classList.add('fadeout');
      }, 800);
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
  transform-style: preserve-3d;
  transition: .8s;

  &_parts {
    position: absolute;
    transition: .5s;
    * {
      transition: .5s;
    }
  }
  &_char {
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
    width: 100%;
  }
  &::after {
    content: '';
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background: none;
    transition: .7s;
  }
  &.white {
    &::after {
      background: #fff;
    }
  }
  &.fadeout {
    opacity: 0;
    pointer-events: none;
  }

  .transform_bg {
    top: 50%;
    transform: translateY(-50%);
  }
  .trans_eyevow {
    width: 100%;
    height: 100%;
    background-repeat: no-repeat;
    background-position: center center;
    background-size: cover;
  }
  #trans_eyevow01 {
    &.illust { background-image: url(../assets/img/eyevow/character_illust_transform.png); }
    &.photo { background-image: url(../assets/img/eyevow/character_photo_transform.png); }
  }
  #trans_eyevow02 {
    &.illust { background-image: url(../assets/img/eyevow/character_illust_02.png); }
    &.photo { background-image: url(../assets/img/eyevow/character_photo_02.png); }
  }
  .transform_char {
    // opacity: 0;
    overflow: hidden;
    height: 0px;

    .img {
      display: block;
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      width: 100%;
      transform: translate(80%, 0%) rotate(20deg);
      filter: blur(10px);
      opacity: 0;
    }
    &.step01 {
      height: 200px;
      .img01 {
        opacity: 1;
        filter: blur(0);
        transform: translate(0, -50%) rotate(0deg);
      }
    }
    &.step02 {
      height: 300px;
      .img01 {
        opacity: 0;
        filter: blur(10px);
        transform: translate(-80%, 0%) rotate(-20deg);
      }
      .img02 {
        opacity: 1;
        filter: blur(0);
        transform: translate(0, -50%) rotate(0deg);
      }
    }
    &.step03 {
      height: 70%;
      .img02 {
        opacity: 0;
        filter: blur(10px);
        transform: translate(-80%, 0%) rotate(-20deg);
      }
      .img03 {
        opacity: 1;
        filter: blur(0);
        transform: translate(0, -50%) rotate(0deg);
      }
    }
    &.step04 {
      height: 90%;
      .img03 {
        opacity: 0;
        filter: blur(10px);
        transform: translate(-80%, 0%) rotate(-20deg);
      }
      .img04 {
        opacity: 1;
        filter: blur(0);
        transform: translate(0, -50%) rotate(0deg);
      }
    }
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
}
</style>
