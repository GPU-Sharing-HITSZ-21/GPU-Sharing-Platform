#!/bin/bash

# 定义一个变量来存储进程 ID，以便出错时可以终止
GO_PID=""
NPM_PID=""

# 启动后端 Go 服务
echo "Starting the Go server..."
cd ./server || { echo "Failed to enter ./server directory"; exit 1; }
nohup go run main.go > ../server.log 2>&1 & 
GO_PID=$!  # 获取 Go 服务的进程 ID
cd .. || { echo "Failed to return to root directory"; kill $GO_PID; exit 1; }

# 启动前端 Vue 应用
echo "Starting the Vue frontend..."
cd ./frontend/GPU-sharing-v2 || { echo "Failed to enter ./frontend/GPU-sharing-v2 directory"; kill $GO_PID; exit 1; }
nohup npm run dev > ../../frontend.log 2>&1 &
NPM_PID=$!  # 获取前端服务的进程 ID

# 返回根目录
cd ../.. || { echo "Failed to return to root directory"; kill $GO_PID $NPM_PID; exit 1; }

# 错误处理：捕获脚本中的任何错误并终止后台进程
trap 'echo "An error occurred. Rolling back..."; kill $GO_PID $NPM_PID; exit 1' ERR

# 显示后台进程
echo "Finished."
