<template>
  <el-container class="container">
    <el-header class="header"></el-header>
    <el-row style="margin:0">
      <el-col :xs="1" :sm="2" :md="3" :lg="3" :xl="4">
        <div class="grid-content"></div>
      </el-col>
      <el-col :xs="20" :sm="20" :md="16" :lg="18" :xl="16">
        <Nav />
        <router-view />
      </el-col>
      <el-col :xs="1" :sm="2" :md="3" :lg="3" :xl="4">
        <div class="grid-content"></div>
      </el-col>
    </el-row>
    <el-footer>
    </el-footer>
  </el-container>
</template>

<script>
  import Nav from './components/common/Nav.vue'

  export default {
    name: 'App',
    components: {
      Nav
    }
  }
</script>

<style>
  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }

  html,
  body,
  #app {
    margin: 0;
    padding: 0;
    height: 100%;
  }

  .grid-content {
    min-height: 36px;
  }

  .container {
    height: 100%;
  }

  .header {
    background-color: rgb(101, 168, 144);
  }

  .aside {
    width: 300px;
    background-color: aquamarine;

  }

  .main {
    background-color: bisque;
    padding: 0px, 100px;
  }
</style>