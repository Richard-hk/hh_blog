import Vue from 'vue'
// TODO: 按需加载组件
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import App from './App.vue';
import './plugin/http'
// import './plugin/elemui'
import router from './router';
Vue.use(ElementUI);
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')