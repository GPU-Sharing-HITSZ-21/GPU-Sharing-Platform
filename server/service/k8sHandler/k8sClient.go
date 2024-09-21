package k8sHandler

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

var K8sClient *kubernetes.Clientset

// 初始化时连接到 Kubernetes
func init() {
	fmt.Println("Initializing Kubernetes client...")
	k8sClientInit()
}

func k8sClientInit() {
	var kubeConfig string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = filepath.Join(home, ".kube", "config")
	}

	// 检查 Kubernetes 配置文件是否存在
	if _, err := os.Stat(kubeConfig); os.IsNotExist(err) {
		fmt.Println("未检测到 Kubernetes 配置，跳过 K8s 客户端初始化")
		return
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		fmt.Println("构建 K8s 配置失败，跳过 K8s 客户端初始化: ", err)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("初始化 K8s 客户端失败: ", err)
		return
	}
	K8sClient = clientset
	fmt.Println("Kubernetes client successfully initialized")
}
