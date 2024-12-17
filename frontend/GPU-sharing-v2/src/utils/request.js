import axios from "axios"
import {ElMessage} from "element-plus"

const request = axios.create({
    baseURL: 'http://10.249.190.219:31024',
    timeout:30000
})

request.interceptors.request.use(config => {
    config.headers['Content-Type'] = 'application/json;charset=utf-8';
    const token = localStorage.getItem('token');
    if (token) {
        config.headers['Authorization'] = `${token}`;
    }
    return config
},error => {
    return Promise.reject(error)
});

request.interceptors.response.use(response => {
    let res = response.data;
    if(typeof res == 'string'){
        res = res ? JSON.parse(res):res
    }
    return res;
    },error => {
    if (error.response.status === 404){
        ElMessage.error("not found")
    }else if(error.response.status === 500){
        ElMessage.error("server exception")
    }else{
        console.error(error.message)
    }
    return Promise.reject(error)
});

export default request