<template>
  <div>
    <!-- <VowPopup /> -->

    <nav id="nav">
      <router-link to="/" class="nav_item nav_home" exact>
        <span class="icon"><iconHome /></span>
        <span class="label">Home</span>
      </router-link>
      <router-link to="/about" class="nav_item nav_about">
        <span class="icon"><iconAbout /></span>
        <span class="label">About</span>
      </router-link>
      <a class="action action_shop" href="https://eyevow.stores.jp/" target="_blank" v-if="isVow">
        <span class="bg"><img alt="" src="@/assets/img/common/icon_action_bg.svg"></span>
        <span class="icon"><iconShop /></span>
        <span class="label">Shop</span>
      </a>
      <router-link class="action action_achieve" to="/achieve" v-else-if="isAchieve"><!-- tether -->
        <span class="bg"><img alt="" src="@/assets/img/common/icon_action_bg.svg"></span>
        <span class="icon"><iconEye /></span>
        <span class="label">Achieved</span>
      </router-link>
      <router-link class="action action_vow" to="/vow" v-else>
        <span class="bg"><img alt="" src="@/assets/img/common/icon_action_bg.svg"></span>
        <span class="icon"><iconEye /></span>
        <span class="label">Vow</span>
      </router-link>
      <router-link to="/cheer" class="nav_item nav_cheer" :class="{'disabled': !isLogin }">
        <span class="icon"><iconCheer /></span>
        <span class="label">Cheer</span>
      </router-link>
      <router-link to="/edit" class="nav_item nav_edit" :class="{'disabled': !isLogin }">
        <span class="icon"><iconEdit /></span>
        <span class="label">Edit</span>
      </router-link>
    </nav>
  </div>
</template>

<script>
// import VowPopup from '@/components/VowPopup.vue'

import iconHome from '@/assets/img/common/icon_home.svg?inline'
import iconAbout from '@/assets/img/common/icon_about.svg?inline'
import iconEye from '@/assets/img/common/icon_eye.svg?inline'
import iconShop from '@/assets/img/common/icon_shop.svg?inline'
import iconCheer from '@/assets/img/common/icon_cheer.svg?inline'
import iconEdit from '@/assets/img/common/icon_edit.svg?inline'

export default {
  name: 'Nav',
  components: {
    // VowPopup,
    iconHome,
    iconAbout,
    iconEye,
    iconShop,
    iconCheer,
    iconEdit
  },
  data: function () {
    return {
      state: {}
    }
  },
  watch: {
    isLogin: function(newVal) {
      return newVal;
    },
    isVow: function(newVal) {
      return newVal;
    },
    isAchieve: function(newVal) {
      return newVal;
    }
  },
  computed: {
    isLogin: function() {
      return this.$store.getters.isLogin
    },
    isVow: function() {
      return this.$store.getters.isVow
    },
    isAchieve: function() {
      return this.$store.getters.isAchieve
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
