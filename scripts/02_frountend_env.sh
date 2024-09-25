sudo yum update
sudo yum install nodejs npm

cd ../frontend/gpu-sharing-platform-frontend || exit

sudo npm install

npm run dev -- --host
