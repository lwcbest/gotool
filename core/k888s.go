package core

import "github.com/lwcbest/gotool/k8s_go"

func RunInK8s(method int) {
	switch method {
	case 1:
		k8s_go.RunRedisDocker()
	case 2:
		k8s_go.DelRedisDocker()
	case 3:
		k8s_go.RunMysqlDocker()
	case 4:
		k8s_go.DelMysqlDocker()
	}
}
