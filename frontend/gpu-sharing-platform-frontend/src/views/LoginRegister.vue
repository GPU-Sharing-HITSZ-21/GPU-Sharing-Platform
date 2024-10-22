<template>
  <div class="auth-container">
    <div class="form-box">
      <h2 class="auth-title">{{ isLogin ? '登录' : '注册' }}</h2>
      <form @submit.prevent="isLogin ? handleLogin() : handleRegister()">
        <div class="input-group">
          <label for="username">用户名</label>
          <input v-model="formData.username" type="text" id="username" required />
        </div>
        <div class="input-group">
          <label for="password">密码</label>
          <input v-model="formData.password" type="password" id="password" required />
        </div>
        <div v-if="!isLogin" class="input-group">
          <label for="role">角色</label>
          <input v-model="formData.role" type="text" id="role" placeholder="默认角色为 USER" />
        </div>
        <button type="submit">{{ isLogin ? '登录' : '注册' }}</button>
      </form>
      <p>
        {{ isLogin ? '没有账户？' : '已有账户？' }}
        <span @click="toggleAuthMode">{{ isLogin ? '注册' : '登录' }}</span>
      </p>
      <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      isLogin: true,
      formData: {
        username: "",
        password: "",
        role: "USER", // 注册时的默认角色
      },
      errorMessage: "",
    };
  },
  methods: {
    toggleAuthMode() {
      this.isLogin = !this.isLogin;
      this.errorMessage = ""; // 切换模式时清除错误消息
    },
    async handleLogin() {
      try {
        const response = await axios.post("/api/user/login", {
          username: this.formData.username,
          password: this.formData.password,
        });
        const {token, role} = response.data; // 提取角色信息
        localStorage.setItem("token", token);
        localStorage.setItem("role", role); // 保存角色信息
        this.$router.push("/pods"); // 登录成功后跳转到仪表盘页面
      } catch (error) {
        this.errorMessage = error.response?.data?.message || "登录失败";
      }
    },
    async handleRegister() {
      try {
        await axios.post("/api/user/register", {
          username: this.formData.username,
          password: this.formData.password,
          role: this.formData.role || "USER", // 使用默认角色
        });
        this.errorMessage = "";
        this.toggleAuthMode(); // 注册成功后自动切换到登录模式
      } catch (error) {
        this.errorMessage = error.response?.data?.error || "注册失败";
      }
    },
  },
};
</script>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.form-box {
  width: 300px;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  text-align: center;
  background-color: rgba(255, 255, 255, 0.9); /* 不透明背景 */
}

.auth-title {
  color: #42b983; /* 更改登录文字的颜色 */
}

.input-group {
  margin-bottom: 20px;
}

input {
  width: 100%;
  padding: 10px;
  margin-top: 5px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #38a877;
}

.error-message {
  color: red;
}

span {
  color: #42b983;
  cursor: pointer;
}
</style>
