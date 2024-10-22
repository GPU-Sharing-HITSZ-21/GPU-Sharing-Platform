package k8sHandler

import (
	"context"
	"github.com/gin-gonic/gin"
	"gpu-sharing-platform/utils"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log" // 导入日志包
	"net/http"
	"regexp"
	"strings"
)

type JobRequest struct {
	Program   string   `json:"program"` // ZIP 文件中可执行文件的名称
	Dataset   []string `json:"dataset"`
	UploadDir string   `json:"uploadDir"`
	InputDir  string   `json:"inputDir"`
	OutputDir string   `json:"outputDir"`
	ZIP       int      `json:"zip"` // 是否为 ZIP 文件
	ZIPName   string   `json:"zipName"`
}

func StartTrainingJob(c *gin.Context) {
	var jobRequest JobRequest

	// 获取请求头中的 JWT token 并解析出用户名
	token := c.Request.Header.Get("Authorization")
	username, err := utils.GetUsername(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: Invalid token"})
		return
	}

	// 解析请求体
	if err := c.ShouldBindJSON(&jobRequest); err != nil {
		log.Printf("请求数据解析失败: %v", err) // 打印错误日志
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	uploadDir := jobRequest.UploadDir + username
	jobName := sanitizeName(jobRequest.Program)
	log.Printf("Job name: %s", jobName)

	var job *batchv1.Job

	// 根据 ZIP 字段决定创建不同的 Job
	if jobRequest.ZIP == 1 {
		// 创建解压并执行的 Job
		job = &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name: jobName,
			},
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  jobName,
								Image: "miniconda-unzip:0.0.1-SNAPSHOT",
								Args: []string{
									"sh", "-c",
									"find /data -name 'lab1.zip' -exec unzip -o {} -d /data \\; && " +
										"cd /data/lab1 && chmod +x " + jobRequest.Program + " && " +
										"python ./" + jobRequest.Program,
								},
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      "data-volume",
										MountPath: jobRequest.InputDir,
									},
									{
										Name:      "output-volume",
										MountPath: jobRequest.OutputDir,
									},
									{
										Name:      "zip-volume",
										MountPath: "/data/zip",
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
										Path: uploadDir,
									},
								},
							},
							{
								Name: "output-volume",
								VolumeSource: corev1.VolumeSource{
									HostPath: &corev1.HostPathVolumeSource{
										Path: "/trainingOpt",
									},
								},
							},
							{
								Name: "zip-volume",
								VolumeSource: corev1.VolumeSource{
									HostPath: &corev1.HostPathVolumeSource{
										Path: uploadDir + "/" + jobRequest.ZIPName, // ZIP 文件路径
									},
								},
							},
						},
						NodeSelector: map[string]string{
							"node-role.kubernetes.io/master": "",
						},
						Tolerations: []corev1.Toleration{
							{
								Key:      "node-role.kubernetes.io/master",
								Operator: corev1.TolerationOpExists,
								Effect:   corev1.TaintEffectNoSchedule,
							},
						},
					},
				},
			},
		}
	} else {
		// 创建直接执行的 Job
		job = &batchv1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name: jobName,
			},
			Spec: batchv1.JobSpec{
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  jobName,
								Image: "miniconda-unzip:0.0.1-SNAPSHOT",
								Args:  []string{"python", "/data/" + jobRequest.Program},
								VolumeMounts: []corev1.VolumeMount{
									{
										Name:      "data-volume",
										MountPath: jobRequest.InputDir,
									},
									{
										Name:      "output-volume",
										MountPath: jobRequest.OutputDir,
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
										Path: uploadDir,
									},
								},
							},
							{
								Name: "output-volume",
								VolumeSource: corev1.VolumeSource{
									HostPath: &corev1.HostPathVolumeSource{
										Path: "/trainingOpt",
									},
								},
							},
						},
						NodeSelector: map[string]string{
							"node-role.kubernetes.io/master": "",
						},
						Tolerations: []corev1.Toleration{
							{
								Key:      "node-role.kubernetes.io/master",
								Operator: corev1.TolerationOpExists,
								Effect:   corev1.TaintEffectNoSchedule,
							},
						},
					},
				},
			},
		}
	}

	// 创建 Job
	_, err = K8sClient.BatchV1().Jobs("default").Create(context.Background(), job, metav1.CreateOptions{})
	if err != nil {
		log.Printf("创建训练任务失败: %v", err) // 打印错误日志
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建训练任务"})
		return
	}

	log.Printf("训练任务已启动: %+v", jobRequest) // 打印成功日志
	c.JSON(http.StatusOK, gin.H{"message": "训练任务已启动", "jobDetails": jobRequest})
}

// sanitizeName 函数将输入名称转换为有效的 Kubernetes 名称
func sanitizeName(name string) string {
	name = strings.TrimSuffix(name, ".py")
	name = strings.TrimSuffix(name, ".zip")
	name = strings.ToLower(name)
	re := regexp.MustCompile("[^a-z0-9-]")
	name = re.ReplaceAllString(name, "-")
	return name
}
