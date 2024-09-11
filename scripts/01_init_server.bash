#!/bin/bash

# 关闭防火墙
echo "正在关闭防火墙..."
systemctl stop firewalld
systemctl disable firewalld
echo "防火墙已关闭。"

# 关闭 SELinux
echo "正在关闭 SELinux..."
sed -i 's/^SELINUX=enforcing/SELINUX=disabled/' /etc/selinux/config  # 永久关闭
setenforce 0  # 临时关闭
echo "SELinux 已关闭。"

# 关闭 swap
echo "正在关闭 swap..."
swapoff -a  # 临时关闭
sed -i '/ swap /s/^/#/' /etc/fstab  # 永久关闭
echo "swap 已关闭。"

# 设置主机名
read -p "请输入新的主机名: " hostname
echo "设置主机名为 $hostname..."
hostnamectl set-hostname $hostname
echo "主机名已设置为 $hostname。"

# 在 master 节点添加 hosts
echo "正在添加 hosts..."
cat >> /etc/hosts << EOF
10.0.12.11 k8s-master
10.0.12.15 k8s-node
EOF
echo "hosts 已添加。"

# 开启内核路由转发
echo "开启内核路由转发..."
sed -i 's/net.ipv4.ip_forward=0/net.ipv4.ip_forward=1/g' /etc/sysctl.conf
sed -i 's/net.ipv4.conf.all.rp_filter=0/net.ipv4.conf.all.rp_filter=1/g' /etc/sysctl.conf

# 将桥接的 IPv4, IPv6 流量传递到 iptables 链
echo "配置桥接的 IPv4 和 IPv6 流量传递到 iptables..."
cat > /etc/sysctl.d/k8s.conf << EOF
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
vm.swappiness = 0
EOF

# 应用配置
sysctl --system
echo "系统配置已应用。"

# 时间同步
echo "安装并执行 ntpdate 时间同步..."
yum install ntpdate -y
ntpdate ntp.ntsc.ac.cn
echo "时间同步完成。"

# 检测是否安装了 Docker
echo "正在检测是否安装 Docker..."
if ! command -v docker &> /dev/null
then
    echo "Docker 未安装，开始下载安装..."
    wget https://mirrors.nju.edu.cn/docker-ce/linux/static/stable/x86_64/docker-20.10.24.tgz
    tar -xf docker-20.10.24.tgz
    cp docker/* /usr/bin
    echo "Docker 已安装。"
else
    echo "Docker 已安装，路径为：$(which docker)"
fi

# 配置 Docker systemd 服务
echo "配置 Docker systemd 服务..."
cat > /etc/systemd/system/docker.service << EOF
[Unit]
Description=Docker Application Container Engine
Documentation=https://docs.docker.com
After=network-online.target firewalld.service
Wants=network-online.target

[Service]
Type=notify
ExecStart=/usr/bin/dockerd
ExecReload=/bin/kill -s HUP \$MAINPID
LimitNOFILE=65535
LimitNPROC=65535
LimitCORE=65535
TimeoutStartSec=0
Delegate=yes
KillMode=process
Restart=on-failure
StartLimitBurst=3
StartLimitInterval=60s

[Install]
WantedBy=multi-user.target
EOF

echo "Docker systemd 服务已配置，正在启动 Docker 服务..."
systemctl daemon-reload
systemctl start docker.service
systemctl enable docker.service
echo "Docker 服务已启动。"

echo "配置镜像加速器"
mkdir -p /etc/docker
tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://w871dh5j.mirror.aliyuncs.com"],
  "exec-opts": ["native.cgroupdriver=systemd"]
}
EOF

systemctl daemon-reload
systemctl restart docker
echo "镜像加速器已配置"

echo "添加阿里镜像源"
cat > /etc/yum.repos.d/kubernetes.repo << EOF
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
echo "镜像源配置完成"

echo "开始安装k8s"
yum install -y kubelet-1.20.0 kubeadm-1.20.0 kubectl-1.20.0
systemctl enable kubelet
echo "安装完成"

echo "开始配置cni"
mkdir -p /etc/cni/net.d
touch /etc/cni/net.d/cni-default.conf
cat > /etc/cni/net.d/cni-default.conf << EOF
{
    "name": "mynet",
    "cniVersion": "0.3.1",
    "type": "bridge",
    "bridge": "mynet0",
    "isDefaultGateway": true,
    "ipMasq": true,
    "hairpinMode": true,
    "ipam": {
        "type": "host-local",
        "subnet": "10.244.0.0/16"
    }
}
EOF
echo "cni配置完成"

echo "所有操作已完成。请检查系统状态。"