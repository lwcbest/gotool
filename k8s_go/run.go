package k8s_go

import (
	"github.com/lwcbest/gotool/k8s_go/mysql_k8s"
	"github.com/lwcbest/gotool/k8s_go/redis_k8s"
)

func RunRedisDocker() {
	svc := &redis_k8s.RedisSVC{}
	k8sManager := &K8SManager{}
	k8sManager.Init(svc)
	k8sManager.StartSvc()
}

func DelRedisDocker() {
	svc := &redis_k8s.RedisSVC{}
	k8sManager := &K8SManager{}
	k8sManager.Init(svc)
	k8sManager.StopSvc()
}

func RunMysqlDocker() {
	svc := &mysql_k8s.MysqlSVC{}
	k8sManager := &K8SManager{}
	k8sManager.Init(svc)
	k8sManager.StartSvc()
}

func DelMysqlDocker() {
	svc := &mysql_k8s.MysqlSVC{}
	k8sManager := &K8SManager{}
	k8sManager.Init(svc)
	k8sManager.StopSvc()
}
