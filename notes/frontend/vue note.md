## vue 使用

### 安装
#### Node.js
18.3 或更高版本的 Node.js https://nodejs.org/zh-cn

1. windows通过安装包进行安装
2. npm 设置国内源 `npm config set registry http://registry.npmmirror.com`
3. `npm install 报错chromedriver status: 404` https://blog.csdn.net/BoReFrontEnd/article/details/105365457 总之安装不上chromedriver,直接把它从package.json中删除后可以正常启动

### 设置前端服务代理

在 `vue.config.js` 中添加代理配置：

```
module.exports = {
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:1024', // 后端 API 地址
        changeOrigin: true,
        pathRewrite: { '^/api': '' }, // 重写路径
      },
    },
  },
};
```
