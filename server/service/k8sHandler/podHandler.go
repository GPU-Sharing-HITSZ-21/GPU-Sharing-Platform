package k8sHandler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/dao"
	"gpu-sharing-platform/models"
	"gpu-sharing-platform/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"sync"
)

// 定义一个全局锁
var podCreationMutex sync.Mutex
var podDeletionMutex sync.Mutex

// CreatePodByUser 创建一个新的 Pod 并为其配置 SSH 服务
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

	// 镜像判断与相应的安装命令
	var installCommand []string
	switch requestBody.Image {
	case "ubuntu":
		requestBody.Image = "ubuntu_ssh:0.0.1-SNAPSHOT"
		installCommand = []string{"/bin/sh", "-c", "service ssh start && tail -f /dev/null"}
	case "centos":
		requestBody.Image = "centos_ssh:0.0.1-Snap-Shot"
		installCommand = []string{"/bin/sh", "-c", "yum install -y openssh-server && /usr/sbin/sshd -D"}
	case "alpine":
		requestBody.Image = "alpine:latest"
		installCommand = []string{"/bin/sh", "-c", "apk add --no-cache openssh && /usr/sbin/sshd -D"}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Image should be ubuntu, centos, or alpine"})
		return
	}

	// 使用全局锁，确保创建 Pod 的互斥
	podCreationMutex.Lock()
	defer podCreationMutex.Unlock()

	// 查询数据库获取最新 podId
	latestPodId, err := dao.GetLatestPodId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to retrieve latest pod ID: %v", err)})
		return
	}

	// 使用用户名生成 Pod 名称，确保唯一性
	podName := fmt.Sprintf("ssh-pod-%d", latestPodId+1) // 假设 podId 从 1 开始

	// 定义 Pod 规范
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   podName,
			Labels: map[string]string{"app": "ssh-pod", "user": username},
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create pod: %v", err)})
		return
	}

	// 为 Pod 创建对应的 SSH Service
	nodeIP, portNum := CreateSshService(pod)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Pod %s created successfully for user %s", podName, username),
	})

	newPod := &models.Pod{
		PodName:    podName,
		Username:   username,
		SSHAddress: nodeIP,
		PortNum:    portNum,
	}

	podID, err := dao.InsertPod(newPod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to insert pod into database: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"podID": podID})
}

// GetPodByUser 根据用户名获取该用户的 Pod 列表，并返回给前端
func GetPodByUser(c *gin.Context) {
	// 从 Authorization header 获取 token
	token := c.Request.Header.Get("Authorization")
	username, err := utils.GetUsername(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid token",
		})
		return
	}

	// 通过用户名获取 Pod 列表
	pods, err := dao.GetPodsByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve pods",
			"error":   err.Error(),
		})
		return
	}

	// 如果没有找到任何 Pod
	if len(pods) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No pods found for the user",
		})
		return
	}

	// 返回 Pod 列表
	c.JSON(http.StatusOK, gin.H{
		"pods": pods,
	})
}

// DeletePodByName 删除指定名称的 Pod 及其对应的 SSH 服务
func DeletePodByName(c *gin.Context) {
	// 获取请求参数中的 Pod 名称
	var requestBody struct {
		PodName string `json:"podName"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil || requestBody.PodName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	podName := requestBody.PodName
	// 使用全局锁，确保删除 Pod 的互斥
	podDeletionMutex.Lock()
	defer podDeletionMutex.Unlock()

	// 删除 Pod
	err := K8sClient.CoreV1().Pods("default").Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to delete pod: %v", err)})
		return
	}

	// 构建 SSH Service 名称
	serviceName := fmt.Sprintf("%s-ssh-service", podName)

	// 删除对应的 SSH Service
	err = K8sClient.CoreV1().Services("default").Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to delete service: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Pod %s and its service %s deleted successfully", podName, serviceName)})
}
