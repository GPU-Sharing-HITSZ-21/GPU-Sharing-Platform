package k8sHandler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func CreateTestPod(c *gin.Context) {
	// 定义 Pod 规范
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "example-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx:latest",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 80,
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
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
