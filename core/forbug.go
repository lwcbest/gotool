package core

import (
	"fmt"
)

func TestForBug() {
	in := []int{1, 2, 3, 4, 5}
	out := make([]*int, 0)
	for _, v := range in {
		//v := v  打开注释即正确
		out = append(out, &v)
	}

	fmt.Println("res:", *out[0], *out[1], *out[2])

	in1 := []int{1, 2, 3, 4, 5}
	out2 := make([]int, 0)
	for _, v := range in1 {
		//v := v  打开注释即正确
		out2 = append(out2, v)
	}

	fmt.Println("res:", out2[0], out2[1], out2[2])
}

func TestForIota(){
	const (
		// NotifyEmail 是发送邮件的通知方式.
		NotifyEmail = 1 << iota
		//NotifySms 是发送短信的通知方式.
		NotifySms
		// NotifyIvr 是发送 ivr 的通知方式.
		NotifyIvr
		// NotifyDChat 是发送 dchat 的通知方式.
		NotifyDChat
		// NotifyAPI 是调用 API 的通知方式.
		NotifyAPI
	)

	const (
		_           = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Println("res:",KB, MB, GB,TB,PB,EB,ZB,YB,498&2)
}