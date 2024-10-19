<template>
  <div class="container">
    <h2>Create Pod</h2>
    <form @submit.prevent="submitForm">
<!--      <div class="form-group">-->
<!--        <label for="cpu">选择 CPU 核数:</label>-->
<!--        <select v-model="selectedCpu" id="cpu">-->
<!--          <option v-for="cpu in cpuOptions" :key="cpu" :value="cpu">{{ cpu }}</option>-->
<!--        </select>-->
<!--      </div>-->

      <div class="form-group">
        <label for="image">选择操作系统:</label>
        <select v-model="selectedSystem" id="image">
          <option v-for="image in systemOptions" :key="image" :value="image">{{ image }}</option>
        </select>
      </div>

      <button type="submit" class="create-button">创建 Pod</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      // selectedCpu: 1, // 默认选择 1 核
      selectedSystem: 'centos', // 默认选择操作系统
      // cpuOptions: [1, 2, 4, 8, 16], // 可选择的 CPU 核数
      systemOptions: ['centos'] // 可选择的操作系统
    };
  },
  methods: {
    async submitForm() {
      try {
        const response = await axios.post('/api/container/create', {
          cpu: this.selectedCpu,
          image: this.selectedSystem, // 传递所选的操作系统
          // 其他必要的参数
        });
        alert('Pod 创建成功！');
        // 跳转到 Pod 列表页面
        this.$router.push('/pods'); // 假设 '/pods' 是 Pod 列表的路由
      } catch (error) {
        console.error("Error creating pod:", error);
        alert('创建 Pod 失败！');
      }
    }
  }
};
</script>

<style>
.container {
  width: 100%; /* 使容器宽度为100% */
  height: 100%; /* 使容器高度为100% */
  padding: 20px; /* 添加内边距 */
  box-sizing: border-box; /* 包含内边距在宽高计算内 */
}

.form-group {
  margin-bottom: 20px; /* 每个表单组之间的间距 */
}

label {
  display: block; /* 标签占一行 */
  margin-bottom: 5px; /* 标签与下方元素的间距 */
}

select {
  width: 100%; /* 下拉框宽度为100% */
  padding: 10px; /* 添加内边距 */
  border: 1px solid #ddd; /* 边框样式 */
  border-radius: 5px; /* 圆角 */
}

.create-button {
  padding: 10px 15px; /* 按钮内边距 */
  background-color: #2c3e50; /* 按钮背景色 */
  color: white; /* 按钮文字颜色 */
  border: none; /* 去掉边框 */
  border-radius: 5px; /* 圆角 */
  cursor: pointer; /* 鼠标指针样式 */
}

.create-button:hover {
  background-color: #34495e; /* 悬停时背景色 */
}
</style>
