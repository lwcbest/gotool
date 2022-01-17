package core

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func TestDataRace(){
	i:=0
	go func() {
		for {
			fmt.Println("i is", i)
			time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
	}
}

func TestDataRaceByMutex(){
	mux :=sync.RWMutex{}
	i:=0
	go func() {
		for {
			mux.RLock()
			fmt.Println("i is", i)
			mux.RUnlock()
			time.Sleep(time.Second)
		}
	}()

	for {
		mux.Lock()
		i += 1
		mux.Unlock()
	}
}

func TestDataRaceByAtomic(){
	i:=int32(0)
	go func() {
		for {
			fmt.Println("i is", atomic.LoadInt32(&i))
			time.Sleep(time.Second)
		}
	}()

	for {
		atomic.AddInt32(&i,1)
	}
}