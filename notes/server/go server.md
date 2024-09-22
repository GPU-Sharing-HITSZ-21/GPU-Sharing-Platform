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

