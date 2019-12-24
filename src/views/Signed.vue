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
    api.post('/signon', {token: q.token})
      .then(res => res.data.token)
      .then(token => localStorage.setItem('USER', token))
      .then(next());
  }
}
</script>
