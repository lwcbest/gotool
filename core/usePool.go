package core

import (
	"fmt"
	jobpool "github.com/lwcbest/gotool/pool"
	"strconv"
	"sync"
	"time"
)

func UsePool(){
	wg := sync.WaitGroup{}
	jp := jobpool.New(3, 50).Start()
	lenth := 100
	wg.Add(lenth)
	for i := 0; i < lenth; i++ {
		fmt.Println("push"+ " "+strconv.Itoa(i))
		jp.PushJobFunc(func(args ...interface{}) {
			defer wg.Done()
			time.Sleep(1*time.Second)
			fmt.Print(args[0].(int), " ")
		}, i)
	}

	wg.Wait()
}