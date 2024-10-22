<template>
  <div>
    <h2>文件上传</h2>
    <input type="file" @change="onFileChange" />
    <button @click="uploadFile">上传</button>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedFile: null,
      message: ''
    };
  },
  methods: {
    onFileChange(event) {
      this.selectedFile = event.target.files[0];
    },
    async uploadFile() {
      if (!this.selectedFile) {
        this.message = '请选择文件';
        return;
      }

      const formData = new FormData();
      formData.append('file', this.selectedFile);

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
