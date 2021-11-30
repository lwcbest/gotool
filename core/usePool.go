package core

import (
	"fmt"
	"github.com/lwcbest/gotool/pool"
	"time"
)

func UsePool(){
	goPool := pool.NewGoPool(2, 1)
	goPool.StartRun()
	for i := 0; i < 106; i++ {
		count := i

		res := goPool.AddTask(func() {
			fmt.Printf("I am task! Number %d\n")
		}, 3*time.Second)

		fmt.Printf("I am save! Number %d  %v\n", count, res)
	}

	// dummy wait until jobs are finished
	time.Sleep(1000 * time.Second)
}