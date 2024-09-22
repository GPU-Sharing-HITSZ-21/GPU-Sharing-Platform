package k8sHandler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"log"
	"net/http"
)

func ExecCommandInPod(clientset *kubernetes.Clientset, config *rest.Config, namespace, podName, containerName string, command []string) (string, error) {
	req := clientset.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		Param("container", containerName).
		Param("command", command[0]).
		Param("stdin", "true").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("tty", "true")

	for _, cmd := range command[1:] {
		req = req.Param("command", cmd)
	}

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return "", err
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    true,
	})
	if err != nil {
		return "", err
	}

	if stderr.Len() > 0 {
		return "", fmt.Errorf(stderr.String())
	}

	return stdout.String(), nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin 用于允许所有的跨域请求（可以根据需要定制）
	CheckOrigin: func(r *http.Request) bool {
		return true // 根据需要修改为只允许特定的 origin
	},
}

func HandleExecWebSocket(c *gin.Context) {
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Failed to upgrade websocket", http.StatusBadRequest)
		return
	}
	defer wsConn.Close()

	restConfig, err := rest.InClusterConfig()
	if err != nil {
		wsConn.WriteMessage(websocket.TextMessage, []byte("Failed to get cluster config"))
		return
	}

	// 这里可以初始化一个 Pod 的执行命令的流
	for {
		// 读取客户端发送的消息
		_, msg, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		// 执行命令
		output, err := ExecCommandInPod(K8sClient, restConfig, "default", "example-pod", "nginx", []string{"/bin/sh", "-c", string(msg)})
		if err != nil {
			wsConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
			continue
		}

		// 发送命令输出回客户端
		err = wsConn.WriteMessage(websocket.TextMessage, []byte(output))
		if err != nil {
			fmt.Println("Error writing WebSocket message:", err)
			break
		}
	}
}
