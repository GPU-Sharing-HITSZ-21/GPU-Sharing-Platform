package k8sHandler

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
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
	command := []string{"/bin/sh"} // 或者 /bin/bash
	output, err := ExecCommandInPod(K8sClient, restConfig, "default", "example-pod", "nginx", command)
	if err != nil {
		err := wsConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
		if err != nil {
			return
		}
		return
	}

	err = wsConn.WriteMessage(websocket.TextMessage, []byte(output))
	if err != nil {
		return
	}
}
