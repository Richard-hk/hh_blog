import Vue from 'vue'
import VueRouter from 'vue-router'
// 页面路由组件
const Index = () => import('../components/index/Index.vue')

const DtDeal = () => import('../components/tool/DtDeal.vue')

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

const routes = [{
  path: '/',
  name: 'index',
  meta: {
    title: '首页'
  },
  component: Index
}, {
  path: '/dtDeal',
  name: 'dtDeal',
  meta: {
    title: '日期处理'
  },
  component: DtDeal
}, ]

const router = new VueRouter({
  routes
})

export default router