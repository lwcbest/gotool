package k8s_go

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type ISvc interface {
	InitMetaData()
	GetName() string
	GetServiceMeta() *corev1.Service
	GetDeploymentMeta() *appsv1.Deployment
	GetConfigMapMeta() *corev1.ConfigMap
}
