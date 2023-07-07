package core

import "github.com/lwcbest/gotool/k8s_go"

func RunRedisInK8s() {
	k8s_go.RunRedisDocker()
}

func DelRedisInK8s() {
	k8s_go.DelRedisDocker()
}
