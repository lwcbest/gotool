package core

import (
	"fmt"
	jobpool "github.com/lwcbest/gotool/pool"
	"time"
)

var jp *jobpool.JobPool

func UsePool() {
	jp = jobpool.New(3, 50).Start()
	newGood := &Good{}
	newGood.sm = SM{state: "1"}
	TestPool(newGood)

	myChan := make(chan int)
	<-myChan
}

type Good struct {
	sm SM
}

type SM struct {
	state string
}

func (sm *SM) SetCurState(s string) {
	sm.state = s
}

func TestPool(good *Good) {
	good.sm.SetCurState(good.sm.state + "1")

	myJob := func(args ...interface{}) {
		time.Sleep(1 * time.Second)
		fmt.Println("this si :" + args[0].(*Good).sm.state + " ")
		TestPool(args[0].(*Good))
	}

	jp.PushJobFunc(myJob, good)
}
