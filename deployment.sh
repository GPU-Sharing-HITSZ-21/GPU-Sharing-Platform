#!/bin/bash

GO_PID=""
NPM_PID=""

echo "Starting the Go server..."
cd ./server || { echo "Failed to enter ./server directory"; exit 1; }
nohup go run main.go > ../server.log 2>&1 & 
GO_PID=$!
cd .. || { echo "Failed to return to root directory"; kill $GO_PID; exit 1; }

echo "Starting the Vue frontend..."
cd ./frontend/GPU-sharing-v2 || { echo "Failed to enter ./frontend/GPU-sharing-v2 directory"; kill $GO_PID; exit 1; }
nohup npm run dev > ../../frontend.log 2>&1 &
NPM_PID=$!

cd ../.. || { echo "Failed to return to root directory"; kill $GO_PID $NPM_PID; exit 1; }

trap 'echo "An error occurred. Rolling back..."; kill $GO_PID $NPM_PID; exit 1' ERR

echo "Finished."
