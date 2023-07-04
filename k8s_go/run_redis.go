package k8s_go

func RunRedisDocker() {
	k8sManager := &K8SManager{}
	k8sManager.Init()
	k8sManager.StartRedis()
}
