<template>
  <div>
    <h2>上传数据集和运行程序</h2>

    <h3>上传数据集</h3>
    <input type="file" @change="onDatasetChange" />

    <h3>上传运行程序</h3>
    <input type="file" @change="onProgramChange" />

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
      selectedDataset: null,
      selectedProgram: null,
      inputDir: '', // 新增 inputDir
      outputDir: ''
    };
  },
  methods: {
    onDatasetChange(event) {
      this.selectedDataset = event.target.files[0];
    },
    onProgramChange(event) {
      this.selectedProgram = event.target.files[0];
    },
    async uploadFilesAndRun() {
      if (!this.selectedDataset || !this.selectedProgram || !this.inputDir || !this.outputDir) {
        alert('请填写所有字段！');
        return;
      }

      const datasetFormData = new FormData();
      datasetFormData.append('file', this.selectedDataset);

      const programFormData = new FormData();
      programFormData.append('file', this.selectedProgram);

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
            program: this.selectedProgram.name, // 运行程序的文件名
            dataset: this.selectedDataset.name, //数据集名称
            uploadDir: this.getUploadDir(), //上传目录
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
    getUploadDir(){
      return '/uploads/'+this.getUsername()
    }
  }
};
</script>
