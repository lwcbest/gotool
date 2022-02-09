package jobpool

import (
	"context"
	"sync"
)

type JobFunc func(args ...interface{})

type Job struct {
	f JobFunc
	args []interface{}
}

type JobPool struct {
	pool        chan *Job
	workerCount int

	stopCtx        context.Context
	stopCancelFunc context.CancelFunc
	wg             sync.WaitGroup
}

func (j *Job) Execute() {
	j.f(j.args...)
}

func New(workerCount, poolLen int) *JobPool {
	return &JobPool{
		workerCount: workerCount,
		pool:        make(chan *Job, poolLen),
	}
}

func (jp *JobPool) PushJob(t *Job) {
	jp.pool <- t
}

func (jp *JobPool) PushJobFunc(f JobFunc,args ...interface{}) {
	jp.pool <- &Job{
		f: f,
		args:args,
	}
}

func (jp *JobPool) work() {
	for {
		select {
		case <-jp.stopCtx.Done():
			jp.wg.Done()
			return
		case t := <-jp.pool:
			t.Execute()
		}
	}
}

func (jp *JobPool) Start() *JobPool {
	jp.wg.Add(jp.workerCount)
	jp.stopCtx, jp.stopCancelFunc = context.WithCancel(context.Background())
	for i := 0; i < jp.workerCount; i++ {
		go jp.work()
	}
	return jp
}

func (jp *JobPool) Stop() {
	jp.stopCancelFunc()
	jp.wg.Wait()
}
