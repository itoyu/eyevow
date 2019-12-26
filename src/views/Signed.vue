<template>
  <div/>
</template>
<script>
import qs from 'qs';
import api from '@/api/client';

export default {
  beforeRouteEnter(to, from, next) {
    const q = qs.parse(location.search, {ignoreQueryPrefix: true});
    console.log(q);

    // Signon
    api.post('/signon', {token: q.token})
      .then(res => res.data.token)
      .then(token => localStorage.setItem('USER', token))
      .then(next());
  },
  beforeCreate() {
    // Set Vow
    api.post('/vows', {
        text: this.$store.getters.myVow.text,
        // type: this.$store.getters.myVow.type
      }, {
      headers: {
        Authorization: `Bearer ${this.$store.state.devData.token}`,
      }
    })
      .then(res => res.data)
      .then(json => {
        // this.$store.commit('setVow');
        console.log(json);

        this.$router.push({ path: '/'})
      })
  }
}
</script>
