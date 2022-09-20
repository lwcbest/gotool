package go_pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func shabile() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("wrong line：%v", v)
		default:
		}
	}
}

func DoStartCPUProfile() {
	file, err := os.Create("./cpu.prof")
	if err != nil {
		fmt.Printf("create file failed, err:%v\n", err)
		return
	}

	_ = pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	for i := 0; i < 4; i++ {
		go shabile()
	}
	time.Sleep(10 * time.Second)
}

func DoListenAndServe() {
	// 执行一段有问题的代码
	for i := 0; i < 4; i++ {
		go shabile()
	}
	http.ListenAndServe("0.0.0.0:6061", nil)
}
