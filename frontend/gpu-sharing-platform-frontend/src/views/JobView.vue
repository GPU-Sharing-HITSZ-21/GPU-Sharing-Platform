<template>
  <div>
    <h2>上传数据集和运行程序</h2>

    <h3>选择上传方式</h3>
    <label>
      <input type="radio" v-model="uploadType" value="separate" />
      单独上传数据集和程序
    </label>
    <label>
      <input type="radio" v-model="uploadType" value="zip" />
      上传压缩包
    </label>

    <div v-if="uploadType === 'separate'">
      <h3>上传数据集</h3>
      <input type="file" multiple @change="onDatasetChange" />

      <h3>上传运行程序</h3>
      <input type="file" @change="onProgramChange" />
    </div>

    <div v-if="uploadType === 'zip'">
      <h3>上传压缩包</h3>
      <input type="file" @change="onZipChange" />

      <h3>程序名称</h3>
      <input v-model="programName" placeholder="请输入程序名称" />

      <h3>ZIP 文件名称</h3> <!-- 新增 ZIP 文件名称输入框 -->
      <input v-model="zipName" placeholder="请输入 ZIP 文件名称" />
    </div>

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
      uploadType: 'separate', // 默认选择单独上传
      selectedDatasets: [],
      selectedProgram: null,
      selectedZip: null,
      programName: '', // 新增程序名称字段
      zipName: '', // 新增 ZIP 文件名称字段
      inputDir: '',
      outputDir: ''
    };
  },
  methods: {
    onDatasetChange(event) {
      this.selectedDatasets = Array.from(event.target.files);
    },
    onProgramChange(event) {
      this.selectedProgram = event.target.files[0];
    },
    onZipChange(event) {
      this.selectedZip = event.target.files[0];
      // 自动填写 zipName
      if (this.selectedZip) {
        this.zipName = this.selectedZip.name; // 设置 ZIP 文件名称
      }
    },
    async uploadFilesAndRun() {
      if (this.uploadType === 'separate') {
        if (this.selectedDatasets.length === 0 || !this.selectedProgram || !this.inputDir || !this.outputDir) {
          alert('请填写所有字段！');
          return;
        }

        // 上传数据集和程序的逻辑
        await this.uploadSeparateFiles();
      } else if (this.uploadType === 'zip') {
        if (!this.selectedZip || !this.programName || !this.zipName || !this.inputDir || !this.outputDir) { // 添加 zipName 检查
          alert('请填写所有字段！');
          return;
        }

        // 上传压缩包的逻辑
        await this.uploadZipFile();
      }
    },
    async uploadSeparateFiles() {
      const datasetFormData = new FormData();
      this.selectedDatasets.forEach(file => {
        datasetFormData.append('files', file);
      });

      const programFormData = new FormData();
      programFormData.append('file', this.selectedProgram);

      try {
        // 上传数据集
        await this.uploadToServer('/api/file/upload', datasetFormData, '数据集上传失败');
        // 上传程序
        await this.uploadToServer('/api/file/upload', programFormData, '运行程序上传失败');

        // 启动训练
        await this.startTraining(false); // false 表示不是压缩包上传
      } catch (error) {
        alert(error.message);
      }
    },
    async uploadZipFile() {
      const zipFormData = new FormData();
      zipFormData.append('file', this.selectedZip);

      try {
        // 上传压缩包
        await this.uploadToServer('/api/file/upload', zipFormData, '压缩包上传失败');

        // 启动训练
        await this.startTraining(true); // true 表示是压缩包上传
      } catch (error) {
        alert(error.message);
      }
    },
    async uploadToServer(url, formData, errorMessage) {
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Authorization': this.getToken()
        },
        body: formData
      });

      if (!response.ok) {
        throw new Error(errorMessage);
      }
    },
    async startTraining(zipUpload = false) {
      const body = {
        program: zipUpload ? this.programName : this.selectedProgram.name, // 使用手动输入的程序名称
        dataset: zipUpload ? [] : this.selectedDatasets.map(file => file.name),
        uploadDir: this.getUploadDir(),
        inputDir: this.inputDir,
        outputDir: this.outputDir,
        ZIP: zipUpload ? 1 : 0, // 添加 ZIP 字段
        zipName: zipUpload ? this.zipName : '' // 添加 zipName 字段
      };

      const jobResponse = await fetch('/api/job/start', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': this.getToken()
        },
        body: JSON.stringify(body)
      });

      if (!jobResponse.ok) {
        throw new Error('启动训练失败');
      }

      alert('数据集和运行程序上传成功，训练程序已启动！');
    },
    getToken() {
      return localStorage.getItem('token');
    },
    getUploadDir() {
      return '/uploads/';
    }
  }
};
</script>
