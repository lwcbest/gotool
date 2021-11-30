package pool

import (
	"fmt"
	"time"
)

func NewGoPool(queueMax int, workerCount int) *GoPool {
	taskQueue := taskQueue{}
	taskQueue.create(queueMax)
	slavePool := slavePool{}
	slavePool.create(workerCount)
	return &GoPool{taskQueue: taskQueue, slavePool: slavePool}
}

type GoPool struct {
	taskQueue taskQueue
	slavePool slavePool
	stop      chan struct{}
}

func (gp *GoPool) AddTask(task func(), wait time.Duration) bool {
	return gp.taskQueue.addTask(task, wait)
}

func (gp *GoPool) StartRun() {
	go func() {
		for {
			select {
			case job := <-gp.taskQueue.tasks:
				slave := <-gp.slavePool.slaves
				slave.job <- job
			case <-gp.stop:
				for i := 0; i < cap(gp.slavePool.slaves); i++ {
					slave := <-gp.slavePool.slaves

					slave.stop <- struct{}{}
					<-slave.stop
				}

				gp.stop <- struct{}{}
				return
			}
		}
	}()
}

func (gp *GoPool) Stop() {
	gp.stop <- struct{}{}
	<-gp.stop
}

type taskQueue struct {
	tasks chan func()
}

func (tq *taskQueue) create(max int) {
	tq.tasks = make(chan func(), max)
}

func (tq *taskQueue) addTask(task func(), wait time.Duration) bool {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(wait)
		timeout <- true
	}()

	select {
	case tq.tasks <- task:
		return true
	case <-timeout:
		return false
	}
}

type slavePool struct {
	slaves chan *slave
}

func (sp *slavePool) create(max int) {
	sp.slaves = make(chan *slave, max)
	for i := 0; i < cap(sp.slaves); i++ {
		sla := slave{parentPool: sp}
		sla.job = make(chan func())
		sla.stop = make(chan struct{})
		sla.start()
	}
}

type slave struct {
	parentPool *slavePool
	job        chan func()
	stop       chan struct{}
}

func (s *slave) start() {
	go func() {
		for {
			// worker free, add it to pool
			s.parentPool.slaves <- s
			select {
			case job := <-s.job:
				fmt.Println("do job")
				job()

			case <-s.stop:
				s.stop <- struct{}{}
				return
			}
		}
	}()
}
