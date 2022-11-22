<template>
  <el-container>
    <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/dtdeal">工具</el-menu-item>
      <el-menu-item index="/notice">通知中心</el-menu-item>
    </el-menu>
  </el-container>
</template>
<script>
  export default {
    data() {
      return {
        activeIndex: '/',
      };
    },
    methods: {
      handleSelect(key) {
        this.$router.push(key).catch((err) => err)
      }
    }
  }
</script>