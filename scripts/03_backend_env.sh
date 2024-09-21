rpm --import https://mirror.go-repo.io/centos/RPM-GPG-KEY-GO-REPO
curl -s https://mirror.go-repo.io/centos/go-repo.repo | tee /etc/yum.repos.d/go-repo.repo
yum install golang

yum install docker-compose

cd ./mysql || exit
docker-compose -f mysql-init.yaml up
cd ..

cd ../server/ || exit

# 设置go install 代理
go env -w GOPROXY=https://goproxy.cn,direct
# 下载依赖
go mod tidy
