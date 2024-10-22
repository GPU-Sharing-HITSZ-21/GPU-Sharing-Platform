<template>
  <div class="container">
    <h2>
      Pod List
      <button @click="createPod" class="create-pod-button">新建 Pod</button>
      <button @click="goToJobView" class="job-view-button">查看作业</button>
      <!-- 仅当角色为 ADMIN 时显示的按钮 -->
      <button v-if="isAdmin" @click="kuboard" class="kuboard">kuboard</button>
      <button v-if="isAdmin" @click="promethues" class="promethues">promethues</button>
    </h2>
    <table>
      <thead>
      <tr>
        <th>ID</th>
        <th>Pod Name</th>
        <th>Username</th>
        <th>SSH Address</th>
        <th>Port Number</th>
        <th>Created At</th>
        <th>Updated At</th>
        <th>Action</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(item, index) in responseList" :key="index">
        <td>{{ item.ID }}</td>
        <td>{{ item.PodName }}</td>
        <td>{{ item.Username }}</td>
        <td>{{ item.ssh_address }}</td>
        <td>{{ item.port_num }}</td>
        <td>{{ item.CreatedAt }}</td>
        <td>{{ item.UpdatedAt }}</td>
        <td>
          <button @click="getSshLink(item)">Get SSH Link</button>
          <button @click="deletePod(item.PodName)">Delete</button>
        </td>
      </tr>
      </tbody>
    </table>

    <div v-if="isModalVisible" class="modal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <h3>SSH Link</h3>
        <p>{{ sshLink }}</p>
        <button class="copy-button" @click="copyToClipboard">Copy to Clipboard</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  data() {
    return {
      responseList: [],
      isModalVisible: false,
      sshLink: '',
      router: null,
      isAdmin: false, // 新增属性判断用户是否为管理员
    };
  },
  mounted() {
    this.router = useRouter();
    this.getResponseData();
    this.checkUserRole(); // 检查用户角色
  },
  methods: {
    async getResponseData() {
      try {
        const response = await axios.get('/api/container/myPods');
        this.responseList = response.data.pods;
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    checkUserRole() {
      const role = localStorage.getItem("role"); // 从 localStorage 获取角色
      this.isAdmin = (role === 'ADMIN'); // 判断是否为管理员
    },
    getSshLink(item) {
      this.sshLink = `ssh root@${item.ssh_address} -p ${item.port_num}`;
      this.isModalVisible = true;
    },
    closeModal() {
      this.isModalVisible = false;
    },
    copyToClipboard() {
      navigator.clipboard.writeText(this.sshLink)
          .then(() => {
            alert('SSH Link copied to clipboard!');
          })
          .catch(err => {
            console.error('Failed to copy: ', err);
          });
    },
    createPod() {
      this.router.push('/create-pod');
    },
    goToJobView() {
      this.router.push('/job');
    },
    async deletePod(podName) {
      const confirmDelete = confirm(`Are you sure you want to delete Pod ${podName}?`);
      if (confirmDelete) {
        try {
          const podToDelete = this.responseList.find(item => item.PodName === podName);
          if (podToDelete) {
            await axios.post(`/api/container/delete`, {
              podName: podToDelete.PodName,
              podId: podToDelete.ID
            });
            alert(`Pod ${podName} deleted successfully!`);
            this.getResponseData();
          } else {
            alert('Pod not found!');
          }
        } catch (error) {
          console.error("Error deleting pod:", error);
          alert('Failed to delete Pod!');
        }
      }
    },
    kuboard() {
      window.open('http://110.40.176.8', '_blank');
    },
    promethues() {
      window.open('http://110.40.176.8:30128/', '_blank');
    },

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

h2 {
  display: flex; /* 使用 Flexbox 布局 */
  justify-content: space-between; /* 两侧对齐 */
  align-items: center; /* 垂直居中对齐 */
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px; /* 与标题之间添加间距 */
}

th, td {
  padding: 10px;
  border: 1px solid #ddd;
  text-align: left;
}

th {
  background-color: #2c3e50;
  color: #ecf0f1;
}

td {
  background-color: #34495e;
  color: #ecf0f1;
}

tbody tr:nth-child(even) {
  background-color: #2c3e50;
}

/* 新建 Pod 按钮样式 */
.create-pod-button,
.job-view-button {
  margin-left: 20px; /* 添加左边距 */
  padding: 5px 10px; /* 添加内边距 */
  background-color: #2c3e50; /* 按钮背景色 */
  color: white; /* 按钮文字颜色 */
  border: none; /* 去掉边框 */
  border-radius: 5px; /* 圆角 */
  cursor: pointer; /* 鼠标指针样式 */
}

.create-pod-button:hover,
.job-view-button:hover {
  background-color: #34495e; /* 悬停时背景色 */
}

/* 模态框样式 */
.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.7); /* 更深的半透明背景 */
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background-color: #fff;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 8px; /* 圆角 */
  width: 90%;
  max-width: 400px; /* 最大宽度 */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* 阴影 */
  display: flex; /* 设置为 Flexbox 布局 */
  flex-direction: column; /* 垂直排列内容 */
  align-items: center; /* 水平居中对齐 */
  justify-content: center; /* 垂直居中对齐 */
}

h3 {
  margin-top: 0; /* 移除标题的上边距 */
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

.copy-button {
  background-color: #2c3e50; /* 按钮背景色 */
  color: white; /* 按钮文字颜色 */
  border: none; /* 去掉边框 */
  padding: 10px 15px; /* 添加内边距 */
  border-radius: 5px; /* 圆角 */
  cursor: pointer; /* 鼠标指针样式 */
  transition: background-color 0.3s; /* 平滑过渡效果 */
  margin-top: 10px; /* 按钮上方添加间距 */
}

.copy-button:hover {
  background-color: #34495e; /* 悬停时背景色 */
}
</style>
