<template>
  <div/>
</template>
<script>
import qs from 'qs';
import api from '@/api/client';


export default {
  // beforeRouteEnter(to, from, next) {
  data: function() {
    return {
      myVow: {
        type: localStorage.getItem('vowType'),
        text: localStorage.getItem('vowText'),
        vowId: ''
      }
    }
  },
  beforeCreate() {
    const q = qs.parse(location.search, {ignoreQueryPrefix: true});
    // console.log(q);

    // # Signon
    api.post('/signon', {token: q.token})
      .then(res => res.data.token)
      .then(token => {
        this.$store.dispatch('setUserToken', { userToken: token })

        if(this.myVow.text !== '') {

          // #Set Vow
          api.post('/vows', {
              text: this.myVow.type,
              type: this.myVow.text
            }, {
            headers: {
              Authorization: `Bearer ${this.$store.state.user.token.userToken}`,
            }
          })
            .then(res => res.data)
            .then(json => {
              console.log(json);
              this.$store.dispatch('initSetVow', {
                vowType: this.myVow.type,
                vowText: this.myVow.text
              }).then(() => {
                localStorage.setItem('vowType', '');
                localStorage.setItem('vowText', '');

                this.$router.push({ path: '/'})
              });
            })
        } else {
          this.$store.commit('isLogin');
          this.$router.push({ path: '/'})
        }
      })
  }
}
</script>
