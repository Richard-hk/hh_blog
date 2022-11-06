import Vue from 'vue'
import axios from 'axios'

// TODO 区分线下和线上配置
// let Url = 'http://localhost:3000/api/v1/'
let Url = 'http://82.156.32.215:3000/api/v1/'

axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export {
  Url
}