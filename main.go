package main

import "github.com/lwcbest/gotool/core"

func main() {
	//core.TestDelMap()
	//core.UsePool()
	//core.TestForBug()
	//core.DoUnsafeTest()
	//core.DoImage()
	//core.TestDataRace()
	//core.TestDataRaceByAtomic()
	//core.ReqStr()
	//core.TestForIota()

	//core.StartServ()
	//core.UseExcel()
	//data_race.DoRace()
	//core.StartReqStrServ()

	//GenerateSql
	rawFileName := "/Users/xxx/book/test/rawData.txt"
	targetFileName := "/Users/xxx/book/test/sqlResult.txt"
	core.GenerateSql(rawFileName, targetFileName)
}
