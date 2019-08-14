<template>
  <div class="cheer">
    <h1>This is an cheer page</h1>
    <p>Cheers</p>
    <ul>
      <li v-for="(item, index) in list" v-bind:key="item.id">
        ID:{{item.id}}
        User:{{item.user.id}}
        Vow:{{item.text}}
        Cheer:{{item.cheer_count}}
        Achieved:{{item.archived}}
        Index: {{index}}
      </li>
    </ul>
  </div>
</template>

<script>
// import axios from 'axios'
import vows from '@/api/vows.json'

export default {
  data: function() {
    return {
      list: []
    }
  },
  created() {
    // axios
    //   .get('./vows.json')
    //   .then(function(res) {
    //     this.list = res.data
    // })
    this.list = vows;
    console.log(this.list);
  }
}
</script>
