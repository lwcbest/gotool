package k8s_go

import (
	"github.com/lwcbest/gotool/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type RedisSVC struct {
	Name       string
	Service    *corev1.Service
	Deployment *appsv1.Deployment
	ConfigMap  *corev1.ConfigMap
}

func (redisSvc *RedisSVC) InitMetaData() {
	redisSvc.Name = "redis"
	redisSvc.createService()
	redisSvc.createDeploy()
	redisSvc.createConfigMap()
}

func (redisSvc *RedisSVC) createService() {
	// 创建Redis服务的定义
	redisService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.Name,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": "redis",
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "redis",
					Port:       6379,
					TargetPort: intstr.FromInt(6379),
					NodePort:   30379,
				},
			},
		},
	}

	redisSvc.Service = redisService
}

func (redisSvc *RedisSVC) createDeploy() {
	// 创建Redis部署的定义
	redisDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "redis",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "redis",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "redis",
							Image: "redis:latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "redis",
									ContainerPort: 6379,
								},
							},
							Args: []string{"--requirepass", "$(REDIS_PASSWORD)"},
							Env: []corev1.EnvVar{
								{
									Name: "REDIS_PASSWORD",
									ValueFrom: &corev1.EnvVarSource{
										ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "redis-config",
											},
											Key: "password",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	redisSvc.Deployment = redisDeployment
}

func (redisSvc *RedisSVC) createConfigMap() {
	// 创建 Redis 密码的 ConfigMap 对象
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.Name + "-config",
		},
		Data: map[string]string{
			"password": "cool123456",
		},
	}

	redisSvc.ConfigMap = configMap
}
