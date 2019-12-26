<template>
  <div/>
</template>
<script>
import qs from 'qs';
import api from '@/api/client';


export default {
  // beforeRouteEnter(to, from, next) {
  beforeCreate() {
    const q = qs.parse(location.search, {ignoreQueryPrefix: true});
    // console.log(q);

    // # Signon
    api.post('/signon', {token: q.token})
      .then(res => res.data.token)
      .then(token => {
        this.$store.dispatch('setUserToken', { userToken: token })

        // #Set Vow
        api.post('/vows', {
            text: this.$store.getters.myVow.text,
            type: this.$store.getters.myVow.type
          }, {
          headers: {
            Authorization: `Bearer ${this.$store.state.user.token.userToken}`,
          }
        })
          .then(res => res.data)
          .then(json => {

            this.$store.dispatch('initSetVow', {
              vowType: this.$store.getters.myVow.type,
              vowText: this.$store.getters.myVow.text
            }).then(() => {
              this.$router.push({ path: '/'})
            });
          })
      })
  }
}
</script>
