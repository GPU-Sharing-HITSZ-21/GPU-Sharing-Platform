<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { computed } from 'vue';
//
// const route = useRoute();
// const isHomeOrLogin = computed(() => {
//   return route.path === '/' || route.path === '/auth';
// });
</script>

<template>
  <div>
    <RouterView />
  </div>
<!--  <div>-->
<!--    &lt;!&ndash; 头部 &ndash;&gt;-->
<!--    <header></header>-->

<!--    <el-container-->
<!--        style="height: 100vh;-->
<!--      border: 1px solid #eee;-->
<!--      margin: 0;-->
<!--      padding: 0;">     &lt;!&ndash; style： 隐藏滚动条  去掉白边&ndash;&gt;-->
<!--      &lt;!&ndash; 侧边栏 &ndash;&gt;-->
<!--      <el-aside style="width: 200px;">-->
<!--        <asideNav></asideNav>-->
<!--      </el-aside>-->
<!--      &lt;!&ndash; 主页面 &ndash;&gt;-->
<!--      <el-main>-->
<!--        <router-view></router-view>-->
<!--      </el-main>-->
<!--    </el-container>-->

<!--  </div>-->
</template>

<style>
</style>
