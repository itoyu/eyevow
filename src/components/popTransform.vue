<template>
  <div>
    <div class="popup popup_confirm">
      <div class="popup_in">
        <dl>
          <dt>
            <p class="chat_item">
              <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
              <span class="txt">じゃあいよいよ目入れするわ。</span>
            </p>
            <p class="chat_item">
              <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
              <span class="txt">∵∴∵∴∵∴∵∴∵∴∵∴<br>■タイプ：{{vowMake.type}}<br>■誓い：{{vowMake.text}}<br>∵∴∵∴∵∴∵∴∵∴∵∴</span>
            </p>
          </dt>
          <dd>
            <nav class="vow_make_nav center"> <!-- v-on:click="vowMakeHandlerCheck" -->
              <router-link  class="btn btn_back" to="/vow">Back</router-link>
              <button @click="startTransform" class="btn btn_vow"><iconEye />Vow !</button>
            </nav>
          </dd>
        </dl>
      </div>
    </div>
  </div>
</template>

<script>
import iconEye from '@/assets/img/icon_eye.svg?inline'

export default {
  data() {
    return {
      vowMake: {
        type: 'illust',
        text: 'test'
      }
    }
  },
  components: {
    iconEye
  },
  created() {
    // this.init();
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
.vow .vow_make_nav.center {
  padding-bottom: 0;
}
</style>
