package k8sHandler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/dao"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"time"
)

func CreateTestPod(c *gin.Context) {
	// 从请求中获取镜像名称
	var requestBody struct {
		Image string `json:"image"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil || requestBody.Image == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// 镜像判断
	var installCommand []string
	switch requestBody.Image {
	case "ubuntu":
		// 使用 ubuntu:20.04 或更高版本，并指定安装 SSH 的命令
		requestBody.Image = "ubuntu_ssh:0.0.1-SNAPSHOT"
		installCommand = []string{
			"/bin/sh", "-c", "service ssh start && tail -f /dev/null",
		}
	case "centos":
		// 使用 centos:7 或 centos:8，并指定安装 SSH 的命令
		requestBody.Image = "centos_ssh:0.0.1-Snap-Shot"
		installCommand = []string{
			"/bin/sh", "-c", "yum install -y openssh-server && /usr/sbin/sshd -D",
		}
	case "alpine":
		// 使用 alpine，并指定安装 SSH 的命令
		requestBody.Image = "alpine:latest"
		installCommand = []string{
			"/bin/sh", "-c", "apk add --no-cache openssh && /usr/sbin/sshd -D",
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image should be ubuntu, centos, or alpine"})
		return
	}

	// 定义 Pod 规范
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "example-pod",
			Labels: map[string]string{"app": "ssh-pod"}, // 添加标签 todo 根据用户创建
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "system",
					Image:   requestBody.Image, // 使用输入的镜像名称
					Command: installCommand,
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 22,
						},
					},
				},
			},
		},
	}

	// 在默认命名空间创建 Pod
	_, err := K8sClient.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create pod: %v", err),
		})
		return
	}

	// 创建对应service
	CreateSshService(pod)

	c.JSON(http.StatusOK, gin.H{})
}

func CreateSshService(createdPod *corev1.Pod) (string, int) {
	// 申请一个合适的端口
	portNum, err := dao.ClaimPort()
	if err != nil {
		return "null", -1 // 申请端口失败
	}

	// 定义 Service 规范
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: createdPod.Name + "-ssh-service", // 使用 Pod 名称作为 Service 名称的一部分
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Port:       22,
					TargetPort: intstr.FromInt(22),
					NodePort:   int32(portNum), // 使用申请的端口
				},
			},
			Selector: map[string]string{"app": "custom-pod", "user": createdPod.Labels["user"],"podName":createdPod.Labels["podName"]}, // 选择标签
		},
	}

	// 在默认命名空间创建 Service
	_, err = K8sClient.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return "null", -1
	}

	// 等待 Pod 运行状态
	for {
		podInfo, err := K8sClient.CoreV1().Pods("default").Get(context.TODO(), createdPod.Name, metav1.GetOptions{})
		if err != nil {
			return "null", -1
		}
		if podInfo.Status.Phase == corev1.PodRunning {
			// 获取节点 IP
			nodeName := podInfo.Spec.NodeName
			node, err := K8sClient.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
			if err != nil {
				return "null", -1
			}

			// 通常选择 InternalIP
			var nodeIP string
			for _, addr := range node.Status.Addresses {
				if addr.Type == corev1.NodeInternalIP {
					nodeIP = addr.Address
					break
				}
			}
			// publicIP, err := dao.GetPublicIpByPrivateIp(nodeIP)
			// if err != nil {
			// 	return "null", -1
			// }
			return nodeIP, portNum
		}

		// 如果 Pod 还没有运行，稍等再检查
		time.Sleep(2 * time.Second)
	}
}
