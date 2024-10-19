import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import axios from "axios";

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')

// 添加请求拦截器
axios.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers['Authorization'] = `${token}`; // 添加 token 到请求头
    }
    return config;
}, error => {
    return Promise.reject(error);
});