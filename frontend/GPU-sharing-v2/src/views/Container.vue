<script setup>
import {onMounted, reactive, ref} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import request from "@/utils/request.js";
import router from "@/router/index.js";
import {CopyDocument} from "@element-plus/icons-vue";

const count = ref(0);
const pods = ref([]);
const load = () => {
  count.value += 2
}

const activeName = ref('1')
const form = reactive({
  name: '',
})

const drawer = ref(false)
const direction = ref('rtl')
const worker = ref('master')
const gpu = ref('0');
const memory = ref(2048);
const diskSize = ref(20);
const cpuCores = ref(2);
const containerName = ref('')

// 获取容器信息
const fetchPods = async () => {
  try {
    const response = await request.get('/container/myPods');
    if (response && response.pods) {
      pods.value = response.pods;
      count.value = response.pods.length;  // 更新容器数量
    }else if(response.message){
      pods.value = [];
      count.value = 0;
      ElMessage.info(response.message);
    }
    else {
      ElMessage.error('Failed to load containers');
    }
  } catch (error) {
    ElMessage.error(`Error: ${error.message}`);
  }
};

// 删除 Pod
const deletePod = async (podId, podName) => {
  try {
    // 确认删除
    await ElMessageBox.confirm(
        `Are you sure you want to delete pod ${podName}?`,
        'Warning',
        {
          confirmButtonText: 'Delete',
          cancelButtonText: 'Cancel',
          type: 'warning',
        }
    );

    // 调用删除接口
    const response = await request.post('/container/delete', {
      podId,
      podName,
    });

    if (response.message) {
      ElMessage.success(response.message);
      await fetchPods(); // 重新加载容器列表
    }
  } catch (error) {
    ElMessage.error(`Failed to delete pod: ${error.message}`);
  }
};

// 组件加载时获取容器数据
onMounted(() => {
  fetchPods();
});

function cancelClick() {
  ElMessageBox.confirm('Are you sure you want to close this?')
      .then(() => {
        drawer.value = false
      })
      .catch(() => {
        // catch error
      })
}

const confirmClick = async () => {
  const containerData = {
    containerName: containerName.value,
    worker: worker.value,
    gpu: gpu.value,
    memory: memory.value,
    diskSize: diskSize.value,
    cpuCores: cpuCores.value
  };

  try {
    // 发送请求创建容器
    const response = await request.post('/container/create', containerData);

    if (response.message) {
      ElMessage.success('Container created successfully!');
      drawer.value = false; // 关闭 Drawer
      await fetchPods();
    } else {
      ElMessage.error('Failed to create container');
    }
  } catch (error) {
    ElMessage.error(`Error: ${error.message}`);
  }
};

const copyToClipboard = (sshAddress) => {
  // Use the Clipboard API to copy the SSH address to the clipboard
  navigator.clipboard.writeText(sshAddress)
      .then(() => {
        ElMessage.success('SSH address copied to clipboard!');
      })
      .catch((error) => {
        ElMessage.error(`Failed to copy: ${error.message}`);
      });
};

</script>

<template>
  <div class="selectivity">
    <el-form class="select-form">
      <el-form-item label="Container name" class="select-form-item">
        <el-input v-model="form.name" style="width: 200px"/>
      </el-form-item>
      <el-form-item class="select-form-item">
        <el-button>Search</el-button>
      </el-form-item>
      <el-form-item>
        <el-button style="background-color: #5384d7;color: white" @click="drawer = true">New Container</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="container-box">
    <el-collapse class="container-list" v-model="activeName" accordion>
      <el-collapse-item v-for="pod in pods" :key="pod.ID" class="container-item" :title="pod.PodName">
      <div class="container-card">
        <div class="container-info-box">
          <div>container info</div>
          <p><strong>Pod Name:</strong> {{ pod.PodName }}</p>
          <p><strong>Username:</strong> {{ pod.Username }}</p>
          <p><strong>SSH Address:</strong>
            <span class="ssh-address">
              ssh root@10.249.190.219 -p {{ pod.port_num }}
              <el-button @click="copyToClipboard(`ssh root@10.249.190.219 -p ${pod.port_num}`)" size="small" style="background-color: rgba(86,146,255,0.61);color: white; margin-left: 10px;">
                <el-icon><CopyDocument /></el-icon>
              </el-button>
            </span>
          </p>
        </div>
        <div class ="container-manage-box">
          <el-button class="container-manage-button">详细信息</el-button>
          <el-button class="container-manage-button" style="background-color: darkred;color: white" @click="deletePod(pod.ID, pod.PodName)">删除容器</el-button>
        </div>
      </div>
      </el-collapse-item>
    </el-collapse>
  </div>
  <div class="page-selector">
    <el-pagination layout="prev, pager, next" :total="count" />
  </div>

  <el-drawer v-model="drawer" :direction="direction">
    <template #header>
      <h2 style="color:#2e86e3;">New Container</h2>
    </template>

    <template #default>
      <el-form label-width="150px" class="form-container">
        <!-- Container Name Input -->
        <el-form-item label="Container Name">
          <el-input v-model="containerName" size="large" placeholder="Enter container name" style="width: 180px"></el-input>
        </el-form-item>

        <el-form-item label="Select Worker">
          <el-radio-group v-model="worker" size="large">
            <el-radio label="master">master</el-radio>
            <el-radio label="node-1">node-1</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="Select GPU">
          <el-radio-group v-model="gpu" size="large">
            <el-radio label="0">GPU0</el-radio>
            <el-radio label="1">GPU1</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item label="Memory Size (MB)">
          <el-input v-model="memory" type="number" :min="1" :max="65536" placeholder="Enter memory size in MB" size="large" style="width: 180px" />
        </el-form-item>

        <el-form-item label="Disk Size (GB)">
          <el-input-number v-model="diskSize" :min="1" :max="1024" label="Disk Size" size="large" />
        </el-form-item>

        <el-form-item label="CPU Cores">
          <el-input-number v-model="cpuCores" :min="1" :max="16" label="CPU Cores" size="large" />
        </el-form-item>
      </el-form>
    </template>

    <template #footer>
      <div class="footer-buttons">
        <el-button @click="cancelClick">Cancel</el-button>
        <el-button type="primary" @click="confirmClick">Confirm</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<style scoped>
.selectivity{
  height: 60px;
  display: flex;
  align-items: center;
  align-content: center;
}

.select-form {
  display: flex;
  align-items: center;
}

.select-form-item{
  margin-right: 10px;
}

.container-list {
  height: calc(100vh - 200px);
  padding: 0;
  margin: 0;
  list-style: none;
}

.container-item{}

.container-box{
  background-color: white;
}

.container-info-box {
  /* 你可以根据需要添加更多样式 */
}

.container-manage-box {
  flex: 0 1 auto;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.container-manage-button{
  width: 120px;
  margin: 0;
}

.container-card{
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  height: 160px;
  background-color: white;
  border-radius: 5px;
  color:black;
  align-content: center;
}

.page-selector{
  display: flex;
  justify-content: center;
}

/* 设置表单容器的对齐方式 */
.form-container {
  display: flex;
  flex-direction: column;
  gap: 20px; /* 控制控件之间的间距 */
}

/* 每一项控件的对齐 */
.el-form-item {
  display: flex;
  align-items: center;
}

/* 对齐标签和输入框 */
.el-form-item .el-form-item__label {
  min-width: 150px; /* 设置标签的宽度 */
  text-align: right; /* 标签右对齐 */
}

/* footer按钮容器 */
.footer-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}
</style>