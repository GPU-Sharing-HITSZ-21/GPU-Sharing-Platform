## server note

### Gin 路由方式

### golang orm框架
GORM + MySQL

1. 安装

    ```bash
    go install gorm.io/gorm@latest
    go install gorm.io/driver/mysql@latest
    ```
    
2. 查看 ``


### kubernetes go

1. 安装

   ```bash
   go get k8s.io/client-go@v0.26.3
   go get k8s.io/api@v0.26.3
   go get k8s.io/apimachinery@v0.26.3
   ```

   

### 测试 websocket

```
curl --include \
	--no-buffer \
	--header "Connection: Upgrade" \
    --header "Upgrade: websocket" \
    --header "Host: localhost:1024" \
    --header "Origin: http://localhost:1024" \
    --header "Sec-WebSocket-Key: SGVsbG8sIHdvcmxkIQ==" \
    --header "Sec-WebSocket-Version: 13" \
    http://localhost:1024/container/terminal
```

```
curl -i -N -H "Connection: Upgrade" -H "Upgrade: websocket" -H "Host: localhost:1024" -H "Origin: localhost:1024" http://localhost:1024/container/terminal
```

### 自定义ssh镜像
centos https://blog.csdn.net/qq_29183811/article/details/121364550
ubuntu 记得加入启动命令
service ssh start && tail -f /dev/null

docker commit -m "centos with ssh" -a "root" centos_ssh centos_ssh:latest
1. 导出 Docker 镜像
   在源机器上，使用 docker save 命令将镜像导出为 tar 文件：

docker save -o <image-file>.tar <image-name>:<tag>
例如，如果您的镜像名为 my-app，标签为 latest，您可以使用：

docker save -o my-app.tar my-app:latest
2. 复制镜像文件
   使用任何文件传输工具（如 scp、rsync 或 USB 驱动器等）将导出的 tar 文件复制到目标机器。例如，如果您使用 scp，可以这样做：

scp my-app.tar user@target-machine:/path/to/destination
3. 导入 Docker 镜像
   在目标机器上，使用 docker load 命令导入 tar 文件中的镜像：

docker load -i <image-file>.tar
例如：

docker load -i my-app.tar
4. 验证导入
   使用以下命令确认镜像是否成功导入：

docker images