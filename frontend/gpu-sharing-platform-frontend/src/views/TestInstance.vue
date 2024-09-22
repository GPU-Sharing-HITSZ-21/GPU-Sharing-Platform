<template>
  <div>
    <h1>Instance List</h1>
    <ul v-if="instances.length">
      <li v-for="instance in instances" :key="instance.Id">
        <strong>{{ instance.Name }}</strong> - Created At: {{ formatDate(instance.CreatedAt) }}
      </li>
    </ul>
    <p v-else>No instances available.</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      instances: []
    };
  },
  created() {
    this.fetchInstances();
  },
  methods: {
    async fetchInstances() {
      try {
        const response = await axios.get('http://localhost:1024/home/get_test_instance');
        this.instances = response.data;
      } catch (error) {
        console.error('Error fetching instances:', error);
      }
    },
    formatDate(dateString) {
      const options = { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit' };
      return new Date(dateString).toLocaleString('en-US', options);
    }
  }
};
</script>

<style>
h1 {
  font-size: 24px;
  color: #333;
  text-align: center;
  margin-bottom: 20px;
}

ul {
  list-style-type: none; /* 去掉默认的列表样式 */
  padding: 0;
  margin: 0;
  max-width: 600px; /* 限制最大宽度 */
  margin-left: auto; /* 居中 */
  margin-right: auto; /* 居中 */
}

li {
  background: #f9f9f9; /* 背景色 */
  border: 1px solid #ddd; /* 边框 */
  border-radius: 5px; /* 圆角 */
  padding: 10px; /* 内边距 */
  margin: 5px 0; /* 外边距 */
  transition: background 0.3s; /* 过渡效果 */
}

li:hover {
  background: #eaeaea; /* 悬停效果 */
}

p {
  text-align: center; /* 居中 */
  color: #777; /* 灰色 */
}
</style>