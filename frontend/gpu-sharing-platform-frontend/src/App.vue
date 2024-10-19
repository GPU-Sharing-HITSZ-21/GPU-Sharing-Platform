<script setup lang="ts">
import { RouterView, useRoute } from 'vue-router';
import { computed } from 'vue';

const route = useRoute();
const isHomeOrLogin = computed(() => {
  return route.path === '/' || route.path === '/auth';
});
</script>

<template>
  <div :class="{ 'background-image': isHomeOrLogin, 'full-screen': !isHomeOrLogin }">
    <RouterView />
  </div>
</template>

<style>
body {
  margin: 0;
  padding: 0;
}

.background-image {
  background-color: #f0f0f0; /* 设置背景颜色 */
  background-image: url('./assets/background.jpg'); /* 设置背景图像（可选） */
  background-size: cover; /* 背景图像填充 */
  background-position: center; /* 背景图像居中 */
  position: fixed; /* 使用固定定位 */
  top: 0;
  left: 0;
  width: 100%; /* 确保宽度为100% */
  height: 100%; /* 确保高度为100% */
}

.full-screen {
  width: 100%;
  height: 100%;
}

header {
  line-height: 1.5;
  max-height: 0;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 0px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;
    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
