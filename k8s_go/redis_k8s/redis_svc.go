package redis_k8s

import (
	"github.com/lwcbest/gotool/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type RedisSVC struct {
	name       string
	service    *corev1.Service
	deployment *appsv1.Deployment
	configMap  *corev1.ConfigMap
}

func (redisSvc *RedisSVC) InitMetaData() {
	redisSvc.name = "redis"
	redisSvc.createService()
	redisSvc.createDeploy()
	redisSvc.createConfigMap()
}

func (redisSvc *RedisSVC) GetName() string {
	return redisSvc.name
}

func (redisSvc *RedisSVC) GetServiceMeta() *corev1.Service {
	return redisSvc.service
}

func (redisSvc *RedisSVC) GetDeploymentMeta() *appsv1.Deployment {
	return redisSvc.deployment
}

func (redisSvc *RedisSVC) GetConfigMapMeta() *corev1.ConfigMap {
	return redisSvc.configMap
}

func (redisSvc *RedisSVC) createService() {
	// 创建Redis服务的定义
	redisService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.name,
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

	redisSvc.service = redisService
}

func (redisSvc *RedisSVC) createDeploy() {
	// 创建Redis部署的定义
	redisDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.name,
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

	redisSvc.deployment = redisDeployment
}

func (redisSvc *RedisSVC) createConfigMap() {
	// 创建 Redis 密码的 ConfigMap 对象
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: redisSvc.name + "-config",
		},
		Data: map[string]string{
			"password": "cool123456",
		},
	}

	redisSvc.configMap = configMap
}
