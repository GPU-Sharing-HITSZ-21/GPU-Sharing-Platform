package k8sHandler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func CreatePodByUser(c *gin.Context) {
	// 获取请求头中的 JWT token 并解析出用户名
	token := c.Request.Header.Get("Authorization")
	username, err := utils.GetUsername(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Invalid token"})
		return
	}

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
		// 使用 ubuntu_ssh:0.0.1-SNAPSHOT 镜像
		requestBody.Image = "ubuntu_ssh:0.0.1-SNAPSHOT"
		installCommand = []string{
			"/bin/sh", "-c", "service ssh start && tail -f /dev/null",
		}
	case "centos":
		// 使用 centos_ssh:0.0.1-Snap-Shot 镜像
		requestBody.Image = "centos_ssh:0.0.1-Snap-Shot"
		installCommand = []string{
			"/bin/sh", "-c", "yum install -y openssh-server && /usr/sbin/sshd -D",
		}
	case "alpine":
		// 使用 alpine 镜像
		requestBody.Image = "alpine:latest"
		installCommand = []string{
			"/bin/sh", "-c", "apk add --no-cache openssh && /usr/sbin/sshd -D",
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image should be ubuntu, centos, or alpine"})
		return
	}

	// 使用用户名生成 Pod 名称，确保唯一性
	podName := fmt.Sprintf("%s-ssh-pod", username)

	// 定义 Pod 规范
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   podName,                                               // 使用用户名生成 Pod 名称
			Labels: map[string]string{"app": "ssh-pod", "user": username}, // 添加用户名标签
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "system",
					Image:   requestBody.Image,
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
	_, err = K8sClient.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to create pod: %v", err),
		})
		return
	}

	// 为 Pod 创建对应的 SSH Service
	CreateSshService(c, pod)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Pod %s created successfully for user %s", podName, username),
	})
}
