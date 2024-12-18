#!/bin/bash

# 查找并杀死 Go 服务进程
GO_PID=$(ps aux | grep 'go run main.go' | grep -v 'grep' | awk '{print $2}')
if [ -z "$GO_PID" ]; then
  echo "Go server is not running."
else
  echo "Stopping Go server with PID $GO_PID..."
  sudo kill $GO_PID
  if [ $? -eq 0 ]; then
    echo "Go server stopped successfully."
  else
    echo "Failed to stop Go server."
  fi
fi

# 查找并杀死 Vue 前端进程
NPM_PID=$(ps aux | grep 'npm run dev' | grep -v 'grep' | awk '{print $2}')
if [ -z "$NPM_PID" ]; then
  echo "Vue frontend is not running."
else
  echo "Stopping Vue frontend with PID $NPM_PID..."
  sudo kill $NPM_PID
  if [ $? -eq 0 ]; then
    echo "Vue frontend stopped successfully."
  else
    echo "Failed to stop Vue frontend."
  fi
fi

# 释放端口 35173
echo "Releasing port 35173..."
sudo fuser -k 35173/tcp
if [ $? -eq 0 ]; then
  echo "Port 35173 released successfully."
else
  echo "Failed to release port 35173."
fi

# 释放端口 31024
echo "Releasing port 31024..."
sudo fuser -k 31024/tcp
if [ $? -eq 0 ]; then
  echo "Port 31024 released successfully."
else
  echo "Failed to release port 31024."
fi

echo "Shutdown process completed."
