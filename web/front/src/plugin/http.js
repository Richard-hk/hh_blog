import Vue from 'vue'
import axios from 'axios'

// TODO 区分线下和线上配置
let Url = process.env.VUE_APP_API_BASEURL

axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export {
  Url
}