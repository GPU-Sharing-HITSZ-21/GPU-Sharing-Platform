package k8sHandler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
		requestBody.Image = "ubuntu:22.04"
		installCommand = []string{
			"/bin/sh", "-c", "service ssh start && tail -f /dev/null",
		}
	case "centos":
		// 使用 centos:7 或 centos:8，并指定安装 SSH 的命令
		requestBody.Image = "centos_ssh:0.0.1-Snap-Shot"
		installCommand = []string{
			"/bin/sh", "-c", "service ssh start && tail -f /dev/null",
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
	CreateSshService(c, pod)

	c.JSON(http.StatusOK, gin.H{})
}

func CreateSshService(c *gin.Context, createdPod *corev1.Pod) {
	// 定义 Service 规范
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ssh-service",
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Ports: []corev1.ServicePort{
				{
					Port:       22,
					TargetPort: intstr.FromInt(22),
					NodePort:   30022, // 指定 NodePort
				},
			},
			Selector: map[string]string{"app": "ssh-pod"}, // 选择标签
		},
	}

	// 在默认命名空间创建 Service
	_, err := K8sClient.CoreV1().Services("default").Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create service: %v", err),
		})
		return
	}

	// 等待 Pod 运行状态
	for {
		podInfo, err := K8sClient.CoreV1().Pods("default").Get(context.TODO(), createdPod.Name, metav1.GetOptions{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Failed to get pod info: %v", err),
			})
			return
		}
		if podInfo.Status.Phase == corev1.PodRunning {
			// 获取节点 IP
			nodeName := podInfo.Spec.NodeName
			node, err := K8sClient.CoreV1().Nodes().Get(context.TODO(), nodeName, metav1.GetOptions{})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Failed to get node info: %v", err),
				})
				return
			}

			// 通常选择 InternalIP
			var nodeIP string
			for _, addr := range node.Status.Addresses {
				if addr.Type == corev1.NodeInternalIP {
					nodeIP = addr.Address
					break
				}
			}

			// 返回 SSH 连接信息
			c.JSON(http.StatusOK, gin.H{
				"ssh": fmt.Sprintf("ssh root@%s -p 30022", nodeIP),
			})
			return
		}

		// 如果 Pod 还没有运行，稍等再检查
		time.Sleep(2 * time.Second)
	}
}
