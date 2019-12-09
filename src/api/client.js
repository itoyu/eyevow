import axios from 'axios';

export default axios.create({
  baseURL: process.env.VUE_APP_API || 'http://eyevow.work.suichu.net/api'
});