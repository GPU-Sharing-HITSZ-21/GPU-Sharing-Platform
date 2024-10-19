#!/bin/bash

# 检查是否有名为 'kuboard' 的容器
if [ $(sudo docker ps -a -q -f name=kuboard) ]; then
    echo "检测到已有名为 'kuboard' 的容器，正在停止并删除..."

    # 停止容器
    sudo docker stop kuboard

    # 删除容器
    sudo docker rm kuboard

    echo "已删除旧的 'kuboard' 容器。"
else
    echo "没有名为 'kuboard' 的容器。"
fi

# 重新创建并启动容器
echo "正在创建并启动新的 'kuboard' 容器..."

sudo docker run -d \
  --restart=unless-stopped \
  --name=kuboard \
  -p 80:80/tcp \
  -p 10081:10081/tcp \
  -e KUBOARD_ENDPOINT="http://10.0.12.11:80" \
  -e KUBOARD_AGENT_SERVER_TCP_PORT="10081" \
  -v /root/kuboard-data:/data \
  swr.cn-east-2.myhuaweicloud.com/kuboard/kuboard:v3

echo "'kuboard' 容器已重新创建并启动。"