package mysql_k8s

import (
	"github.com/lwcbest/gotool/utils"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type MysqlSVC struct {
	name       string
	service    *corev1.Service
	deployment *appsv1.Deployment
	configMap  *corev1.ConfigMap
}

func (svc *MysqlSVC) InitMetaData() {
	svc.name = "mysql"
	svc.createService()
	svc.createDeploy()
	svc.createConfigMap()
}

func (svc *MysqlSVC) GetName() string {
	return svc.name
}

func (svc *MysqlSVC) GetServiceMeta() *corev1.Service {
	return svc.service
}

func (svc *MysqlSVC) GetDeploymentMeta() *appsv1.Deployment {
	return svc.deployment
}

func (svc *MysqlSVC) GetConfigMapMeta() *corev1.ConfigMap {
	return svc.configMap
}

func (svc *MysqlSVC) createService() {
	// 创建Redis服务的定义
	k8sSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: svc.name,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": "mysql",
			},
			Ports: []corev1.ServicePort{
				{
					Name:       "mysql",
					Port:       3306,
					TargetPort: intstr.FromInt(3306),
					NodePort:   30306,
				},
			},
		},
	}

	svc.service = k8sSvc
}

func (svc *MysqlSVC) createDeploy() {
	// 创建Redis部署的定义
	k8sDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: svc.name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: utils.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "mysql",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "mysql",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "mysql",
							Image: "mysql:latest",
							Ports: []corev1.ContainerPort{
								{
									Name:          "mysql",
									ContainerPort: 3306,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name: "MYSQL_ROOT_PASSWORD",
									ValueFrom: &corev1.EnvVarSource{
										ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "mysql-config",
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

	svc.deployment = k8sDeployment
}

func (svc *MysqlSVC) createConfigMap() {
	// 创建 Redis 密码的 ConfigMap 对象
	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name: svc.name + "-config",
		},
		Data: map[string]string{
			"password": "cool123456",
		},
	}

	svc.configMap = configMap
}
