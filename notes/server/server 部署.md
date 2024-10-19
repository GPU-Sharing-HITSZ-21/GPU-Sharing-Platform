#### 报这个failed: open /run/flannel/subnet.env: no such file or directory
新建 /run/flannel/subnet.env 这个文件写入内容：

FLANNEL_NETWORK=10.244.0.0/16
FLANNEL_SUBNET=10.244.0.1/24
FLANNEL_MTU=1450
FLANNEL_IPMASQ=true