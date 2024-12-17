<script setup>
import router from "@/router/index.js";
</script>

<template>
  <div>
    <!-- Header -->
    <div style="height: 60px;background-color: #5384d7;display: flex;align-items: center">
      <div style="width: 200px;margin-left: 10px;display: flex;align-items: center;font-size: 20px;color: white">
        GPU-Sharing
      </div>
      <div style="flex: 1"></div>
      <div style="width: fit-content;display: flex;align-items: center;padding-right: 20px">Avatar</div>
    </div>

    <!-- Content -->
    <div style="display: flex">
      <!-- Navigation -->
      <div style="width: 200px;border-right: 1px solid #b8b8b8;min-height: calc(100vh - 60px)">
        <el-menu router :default-active="router.currentRoute.value.path" style="border: 0">
          <el-menu-item index="/manager/dashboard">首页</el-menu-item>
          <el-menu-item index="/manager/container">容器</el-menu-item>
        </el-menu>
      </div>
      <!-- Dashboard -->
      <div style="flex: 1;width: 0;background-color: white;padding: 10px">
        <RouterView/>
      </div>
    </div>
  </div>
</template>

<style scoped>
.el-menu .is-active{
  background-color: #3e5a9e;
  color: white;
}
</style>