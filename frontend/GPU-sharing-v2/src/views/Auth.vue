<script setup>
  import {reactive} from "vue";
  import {User,Lock} from "@element-plus/icons-vue"
  import router from "@/router/index.js";
  import request from "@/utils/request.js";

  const data = reactive({
    form:{}
  })

  const handleLogin = () => {
    const { username, password } = data.form;
    request.post('/user/login',{
      username:username,
      password:password
    })
        //todo: 这个返回值封装不行
        .then(response => {
          console.log(response)
          const { role, token } = response;
          if (role === 'ADMIN') {

            // 存储 token 和 role
            localStorage.setItem('token', token);
            localStorage.setItem('role', role);

            // 登录成功后根据角色跳转页面
            if (role === 'ADMIN') {
              router.push({ path: '/manager/dashboard' });
            }
          }
        })
        .catch(error => {
          console.error('Error:', error);
        });

  };
</script>

<template>
  <div class="auth-container">
    <div class="auth-box">
      <div style="padding: 20px;background-color: white;margin-left: 140px;border-radius: 5px">
        <el-form ref="formRef" :model="data.form" style="width: 400px">
          <div style="margin-bottom: 20px;font-size: 24px;text-align: center;color: #3e5a9e;font-weight: bold">Login</div>
          <el-form-item>
            <el-input size="large" v-model="data.form.username" placeholder="username" prefix-icon="User"></el-input>
          </el-form-item>
          <el-form-item>
            <el-input size="large" v-model="data.form.password" placeholder="password" prefix-icon="Lock"></el-input>
          </el-form-item>

          <div>
            <el-button style="width: 47%" @click="handleLogin">Login</el-button>
            <el-button style="width: 47%;float: right">Register</el-button>
          </div>
        </el-form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.auth-container{
  height: 100vh;
  overflow: hidden;
  background-color: #3e5a9e;
  //background-image: ;
}

.auth-box{
  position: absolute;
  display: flex;
  align-items: center;
  width: 50%;
  right: 0;
  height: 100%;
}
</style>