<template>
  <div class="vow_make" id="vow_make">
    <!-- Step 01 -->
    <dl class="vow_make_item step01">
      <dt>
        <span class="dt_step">STEP 1</span>
        <span class="dt_txt">eyevowのタイプを選んでください。</span>
      </dt>
      <dd>
        <div class="vow_make_type">
          <ul class="vow_make_type_slide">
            <li>
              <figure class="img"><img alt="" src="@/assets/img/eyevow/character_illust_02.png"></figure>
              <p class="txt">イラスト</p>
            </li>
            <li>
              <figure class="img"><img alt="" src="@/assets/img/eyevow/character_cosplayers_02.png"></figure>
              <p class="txt">レイヤー</p>
            </li>
            <li>
              <figure class="img soon">Coming Soon.</figure>
              <p class="txt">フィギュア</p>
            </li>
          </ul>
          <button class="btn">Select</button>
        </div>
      </dd>
    </dl>

    <!-- Step 02 -->
    <dl class="vow_make_item step02">
      <dt>
        <span class="dt_step">STEP 2</span>
        <span class="dt_txt">誓いを入力してください。</span>
      </dt>
      <dd>
        <div class="vow_make_input">
          <textarea name="vow_input" rows="4" cols="80" placeholder="誓いや目標をここに入力して、このeyevowに目入れしましょう"></textarea>
          <p class="note">※誓いは後から変更できません</p>
          <button class="btn">Next</button>
        </div>
      </dd>
    </dl>

    <!-- Step 03 -->
    <dl class="vow_make_item step03">
      <dt>
        <span class="dt_step">STEP 3</span>
        <span class="dt_txt">誓う</span>
      </dt>
      <dd>
        <div class="vow_make_input" style="margin-bottom: 100px;">
          <p class="txt txt_vow">宇宙飛行士になるぞ！</p>
          <p class="txt">この内容でeyevowに誓いますか？</p>
          <p class="note">※誓いは後から変更できません</p>
          <router-link class="btn btn_vow" to="/eyevow">VOW</router-link>
        </div>
      </dd>
    </dl>
  </div>
</template>


<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
