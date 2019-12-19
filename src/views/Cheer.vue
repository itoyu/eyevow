<template>
  <div class="cheer">
    <div class="page_lead">
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">みんなの誓いがここに集まっているわ</span>
      </p>
      <p class="chat_item">
        <span class="icon"><img alt="profile icon" src="@/assets/img/eyevow/icon_illust_01.png"></span>
        <span class="txt">応援したいものがあったら「Cheer」ボタンを押して声援を送ってあげて</span>
      </p>
    </div>

    <ul class="tab_nav">
      <li @click="tabSort('all')" class="tab_nav_item all active">Hot</li>
      <li @click="tabSort('progress')" class="tab_nav_item progress">Progress</li>
      <li @click="tabSort('achived')" class="tab_nav_item achived">Achived</li>
    </ul>

    <div class="tab_contents">
      <div class="tab_contents_item">
        <ul class="cheer_list">
          <!-- {{vowlist}} -->
          <li v-for="(vows) in vowlist" v-bind:key="vows.id" v-bind:id="vows.id" class="cheer_list_item show" v-bind:class="{ achieved: vows.archived }">
            <span class="icon_achieve" v-show="vows.archived"><iconAchieve /></span>
            <figure class="cheer_icon"><img alt="" v-bind:src="vows.user.icon"></figure>
            <div class="cheer_info">
              <p class="cheer_name">{{vows.user.name}}</p>
              <p class="cheer_text">{{vows.text}}</p>
              <p class="cheer_time">TIME</p>
            </div>
            <div class="cheer_action">
              <p class="cheer_action_num">{{vows.cheer_count}}</p>
              <button class="cheer_action_btn" @click="vowCheer(vows.id, $event)"><iconLike /></button>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import api from '@/api/client';
// import vows from '@/api/vows.json'
import iconLike from '@/assets/img/icon_like.svg?inline'
import iconAchieve from '@/assets/img/icon_achieve.svg?inline'

export default {
  name: 'Cheer',
  data: function() {
    return {
      vowlist: [],
      token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiNWRmNGYzZDk2YTM0MzMxNGNlODUxNDM2IiwiZXhwIjoxNzM0MDE0Mjk3fQ.37yVyjK5fR9JVc3MOgqZUkpmDJlLTDQ61gPSWFIs1-o'
    }
  },
  components: {
    iconLike,
    iconAchieve
  },
  methods: {
    vowCheer: function(vowid) {
      const cheerBtn = document.getElementById(vowid).getElementsByClassName('cheer_action_btn');

      if(cheerBtn[0].classList.contains('cheered')) {
        api.delete(`vows/${vowid}/cheer`, {
          headers: {
            Authorization: `Bearer ${this.token}`,
          }
        })
          // .then(res => {
          .then(function() {
            cheerBtn[0].classList.remove('cheered');
            console.log('remove cheered');
          })
      } else {
        api.put(`vows/${vowid}/cheer`, {}, {
          headers: {
            Authorization: `Bearer ${this.token}`,
          }
        })
          .then(function() {
            cheerBtn[0].classList.add('cheered');
            console.log('add cheered');
          })
      }
    },
    tabSort: function(sort) {
      var cheersSort;
      const cheersAll = document.querySelectorAll('.cheer_list_item'),
            tabAll = document.querySelectorAll('.tab_nav_item');

      cheersAll.forEach((el) => {
        el.classList.remove('show');
      });
      tabAll.forEach((el) => {
        el.classList.remove('active');
      });

      if(sort === 'progress') {
        cheersSort = document.querySelectorAll('.cheer_list_item:not(.achieved)');
      } else if(sort === 'achived') {
        cheersSort = document.querySelectorAll('.cheer_list_item.achieved');
      } else {
        cheersSort = document.querySelectorAll('.cheer_list_item');
      }

      document.querySelector(`.tab_nav_item.${sort}`).classList.add('active');

      cheersSort.forEach((el) => {
        el.classList.add('show');
      });
    }
  },
  mounted () {
    api.get('/vows')
      .then(res => res.data)
      .then(json => {
        this.vowlist = json.vows.reverse();
      })
  }
}
</script>
