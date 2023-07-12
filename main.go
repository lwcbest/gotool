package main

import (
	"flag"
	"fmt"

	"github.com/lwcbest/gotool/core"
)

func main() {
	cmd := ""
	flag.StringVar(&cmd, "c", "", "要执行的命令，默认为空。")
	flag.Parse()

	switch cmd {
	case "1":
		core.DoLC()
	case "2":
		core.UsePool()
	case "3":
		core.DoGenSql()
	case "4":
		core.DoImage()
	case "run_redis":
		fmt.Println("run_redis_start")
		core.RunInK8s(1)
	case "del_redis":
		fmt.Println("del_redis_start")
		core.RunInK8s(2)
	case "run_mysql":
		fmt.Println("run_mysql_start")
		core.RunInK8s(3)
	case "del_mysql":
		fmt.Println("del_mysql_start")
		core.RunInK8s(4)
	default:
		fmt.Println("abc")
	}

	//core.TestForBug()
	//core.DoUnsafeTest()
	//core.TestDataRace()
	//core.TestDataRaceByAtomic()
	//core.ReqStr()
	//core.TestForIota()
	//core.StartServ()
	//core.UseExcel()
	//data_race.DoRace()
	//core.StartReqStrServ()
}
