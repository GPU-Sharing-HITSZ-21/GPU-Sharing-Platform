package k8sHandler

import (
	"context"
	"github.com/gin-gonic/gin"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log" // 导入日志包
	"net/http"
)

type JobRequest struct {
	Program   string `json:"program"`
	Dataset   string `json:"dataset"`
	UploadDir string `json:"uploadDir"`
	InputDir  string `json:"inputDir"`
	OutputDir string `json:"outputDir"`
}

func StartTrainingJob(c *gin.Context) {
	var jobRequest JobRequest

	// 解析请求体
	if err := c.ShouldBindJSON(&jobRequest); err != nil {
		log.Printf("请求数据解析失败: %v", err) // 打印错误日志
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 创建 Job 对象
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobRequest.Program, // 使用程序名称作为 Job 名称
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  jobRequest.Program,
							Image: "continuumio/miniconda3", // 容器镜像
							Args:  []string{},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "data-volume",       // Volume 名称
									MountPath: jobRequest.InputDir, // 容器内路径
								},
								{
									Name:      "output-volume",      // opt 名称
									MountPath: jobRequest.OutputDir, // 容器内路径
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{
							Name: "data-volume",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: jobRequest.UploadDir, // 上传的目录路径
								},
							},
						},
						{
							Name: "output-volume",
							VolumeSource: corev1.VolumeSource{
								HostPath: &corev1.HostPathVolumeSource{
									Path: "/trainingOpt", // 输出目录路径
								},
							},
						},
					},
				},
			},
		},
	}

	// 创建 Job
	_, err := K8sClient.BatchV1().Jobs("default").Create(context.Background(), job, metav1.CreateOptions{})
	if err != nil {
		log.Printf("创建训练任务失败: %v", err) // 打印错误日志
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建训练任务"})
		return
	}

	log.Printf("训练任务已启动: %+v", jobRequest) // 打印成功日志
	c.JSON(http.StatusOK, gin.H{"message": "训练任务已启动", "jobDetails": jobRequest})
}
