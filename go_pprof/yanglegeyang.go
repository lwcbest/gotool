package go_pprof

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func DoYangLeGeYang() {
	AllYangs := [4]Yang{
		&LieYangYang{},
		&NuanYangYang{},
		&LanYangYang{},
		&MeiYangYang{},
	}

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		if err := http.ListenAndServe(":6061", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		for _, v := range AllYangs {
			v.Go()
		}
		log.Println("----")
		time.Sleep(time.Second)
	}
}

type Yang interface {
	Name() string
	Go()
	Run()
	Jump()
	Fight()
	Idle()
}

type LieYangYang struct {
}

func (y *LieYangYang) Name() string {
	return "烈羊羊"
}

func (y *LieYangYang) Go() {
	y.Run()
	y.Jump()
	y.Fight()
	y.Idle()
}

func (y *LieYangYang) Run() {
	log.Println(y.Name(), "Run")
}

func (y *LieYangYang) Jump() {
	log.Println(y.Name(), "Jump")
}

func (y *LieYangYang) Fight() {
	log.Println(y.Name(), "Fight")
	//打架太厉害了
	//_ = make([]byte, 10*1024*1024)
}

func (y *LieYangYang) Idle() {
	log.Println(y.Name(), "Idle")
}

type NuanYangYang struct {
}

func (y *NuanYangYang) Name() string {
	return "暖羊羊"
}

func (y *NuanYangYang) Go() {
	y.Run()
	y.Jump()
	y.Fight()
	y.Idle()
}

func (y *NuanYangYang) Run() {
	log.Println(y.Name(), "Run")
}

func (y *NuanYangYang) Jump() {
	log.Println(y.Name(), "Jump")
}

func (y *NuanYangYang) Fight() {
	log.Println(y.Name(), "Fight")
	//不想打架
	//m := &sync.Mutex{}
	//m.Lock()
	//go func() {
	//	time.Sleep(time.Second)
	//	m.Unlock()
	//}()
	//m.Lock()
}

func (y *NuanYangYang) Idle() {
	log.Println(y.Name(), "Idle")
	//平静一下
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		time.Sleep(30 * time.Second)
	//	}()
	//}
}

type LanYangYang struct {
}

func (y *LanYangYang) Name() string {
	return "懒羊羊"
}

func (y *LanYangYang) Go() {
	y.Run()
	y.Jump()
	y.Fight()
	y.Idle()
}

func (y *LanYangYang) Run() {
	log.Println(y.Name(), "Run")
}

func (y *LanYangYang) Jump() {
	log.Println(y.Name(), "Jump")
	// 不想跳
	//times := 10000000000
	//for i := 0; i < times; i++ {
	//	// nothing
	//}
}

func (y *LanYangYang) Fight() {
	log.Println(y.Name(), "Fight")
}

func (y *LanYangYang) Idle() {
	log.Println(y.Name(), "Idle")
	// 躺一会
	<-time.After(time.Second)
}

type MeiYangYang struct {
	buffer []byte
}

func (y *MeiYangYang) Name() string {
	return "美羊羊"
}

func (y *MeiYangYang) Go() {
	y.Run()
	y.Jump()
	y.Fight()
	y.Idle()
}

func (y *MeiYangYang) Run() {
	log.Println(y.Name(), "Run")
	//var max int64 = 10 * 1024 * 1024
	//for int64(len(y.buffer)) < max {
	//	y.buffer = append(y.buffer, 'a')
	//}
}

func (y *MeiYangYang) Jump() {
	log.Println(y.Name(), "Jump")
}

func (y *MeiYangYang) Fight() {
	log.Println(y.Name(), "Fight")
}

func (y *MeiYangYang) Idle() {
	log.Println(y.Name(), "Idle")
}
