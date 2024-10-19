<template>
  <div class="container">
    <h1 class="title">GPU 共享平台</h1>
    <div class="button-container">
      <router-link to="/auth" class="button">登录</router-link>
    </div>
  </div>
</template>

<script setup lang="ts">
// 这里可以放置你的逻辑代码
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;700&display=swap');

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh; /* 使内容垂直居中 */
}

.title {
  font-family: 'Roboto', sans-serif;
  color: rgba(0, 128, 0, 0.6); /* 低对比度绿色 */
  font-size: 2.5rem; /* 根据需要调整标题大小 */
  font-weight: 500; /* 调整字体粗细 */
}

.button-container {
  margin-top: 20px; /* 按钮与标题的间距 */
}

.button {
  margin: 0 10px; /* 按钮之间的间距 */
  padding: 10px 20px;
  background-color: #4CAF50; /* 按钮背景颜色 */
  color: white; /* 按钮文字颜色 */
  border: none;
  border-radius: 5px; /* 按钮圆角 */
  text-decoration: none; /* 去掉链接下划线 */
  transition: background-color 0.3s; /* 添加过渡效果 */
  font-family: 'Roboto', sans-serif; /* 按钮字体 */
  font-weight: 500; /* 按钮字体粗细 */
}

.button:hover {
  background-color: #45a049; /* 悬停时的背景颜色 */
}
</style>