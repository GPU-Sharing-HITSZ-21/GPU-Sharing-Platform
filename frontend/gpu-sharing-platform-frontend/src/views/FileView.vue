<template>
  <div>
    <h2>文件上传</h2>
    <input type="file" @change="onFileChange" multiple />
    <button @click="uploadFiles">上传</button>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedFiles: [], // 修改为数组以存储多个文件
      message: ''
    };
  },
  methods: {
    onFileChange(event) {
      this.selectedFiles = Array.from(event.target.files); // 转换为数组
    },
    async uploadFiles() {
      if (this.selectedFiles.length === 0) {
        this.message = '请选择文件';
        return;
      }

      const formData = new FormData();
      this.selectedFiles.forEach(file => {
        formData.append('files[]', file); // 使用数组形式添加文件
      });

      try {
        const response = await fetch('/api/file/upload', {
          method: 'POST',
          body: formData
        });
        if (response.ok) {
          this.message = '文件上传成功';
        } else {
          this.message = '文件上传失败';
        }
      } catch (error) {
        this.message = '发生错误：' + error.message;
      }
    }
  }
};
</script>

<style scoped>
/* 样式可根据需要自定义 */
</style>
