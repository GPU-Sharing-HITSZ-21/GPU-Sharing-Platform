<template>
  <div>
    <h2>上传数据集和运行程序</h2>

    <h3>上传数据集</h3>
    <input type="file" multiple @change="onDatasetChange" />

    <h3>上传运行程序</h3>
    <input type="file" @change="onProgramChange" /> <!-- 只允许上传一个程序 -->

    <h3>输入目录</h3>
    <input v-model="inputDir" placeholder="输入目录" />

    <h3>输出目录</h3>
    <input v-model="outputDir" placeholder="输出目录" />

    <button @click="uploadFilesAndRun">上传并运行</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedDatasets: [], // 数组以支持多个数据集
      selectedProgram: null, // 单个程序
      inputDir: '',
      outputDir: ''
    };
  },
  methods: {
    onDatasetChange(event) {
      this.selectedDatasets = Array.from(event.target.files); // 转换为数组
    },
    onProgramChange(event) {
      this.selectedProgram = event.target.files[0]; // 只获取第一个文件
    },
    async uploadFilesAndRun() {
      if (this.selectedDatasets.length === 0 || !this.selectedProgram || !this.inputDir || !this.outputDir) {
        alert('请填写所有字段！');
        return;
      }

      // 创建 FormData 对象
      const datasetFormData = new FormData();
      this.selectedDatasets.forEach(file => {
        datasetFormData.append('files', file); // 使用 'files' 作为字段名
      });

      const programFormData = new FormData();
      programFormData.append('file', this.selectedProgram); // 只上传一个程序

      try {
        // 上传数据集
        const datasetUploadResponse = await fetch('/api/file/upload', {
          method: 'POST',
          headers: {
            'Authorization': this.getToken()
          },
          body: datasetFormData
        });

        if (!datasetUploadResponse.ok) {
          throw new Error('数据集上传失败');
        }

        // 上传运行程序
        const programUploadResponse = await fetch('/api/file/upload', {
          method: 'POST',
          headers: {
            'Authorization': this.getToken()
          },
          body: programFormData
        });

        if (!programUploadResponse.ok) {
          throw new Error('运行程序上传失败');
        }

        // 调用 API 启动训练
        const jobResponse = await fetch('/api/job/start', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': this.getToken()
          },
          body: JSON.stringify({
            program: this.selectedProgram.name, // 只获取一个程序的文件名
            dataset: this.selectedDatasets.map(file => file.name), // 数据集名称数组
            uploadDir: this.getUploadDir(), // 上传目录
            inputDir: this.inputDir, // 数据集存储路径
            outputDir: this.outputDir // 输出目录
          })
        });

        if (!jobResponse.ok) {
          throw new Error('启动训练失败');
        }

        alert('数据集和运行程序上传成功，训练程序已启动！');
      } catch (error) {
        alert(error.message);
      }
    },
    getToken() {
      return localStorage.getItem('token');
    },
    getUsername() {
      return localStorage.getItem('username');
    },
    getUploadDir() {
      return '/uploads/';
    }
  }
};
</script>
