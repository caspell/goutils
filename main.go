package main

import (
	"fmt"
	"time"

	_ "github.com/wunicorns/goutils/batch"
	_ "github.com/wunicorns/goutils/httprequest"
	"github.com/wunicorns/goutils/patterns"
	_ "github.com/wunicorns/goutils/querybuilder"
)

func main() {

	// fmt.Println("start")

	// filename := "querybuilder/sample.yml"

	// querybuilder.QueryBuilderInitialize(filename)

	// stmt := querybuilder.GetQueryBuilder().GetStatement("get_sample_list")

	// fmt.Println(stmt.Script)

	// fmt.Println("start")

	// api := httprequest.RestAPI{
	// 	Host: "https://192.168.0.8:30000",
	// }

	// if str, err := api.GetJson("/rest/v1/devices/0/diagnostics", nil); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(str)
	// }

	for i, _ := range make([]int, 10) {
		go patterns.Execute(fmt.Sprintf("id: %d", i))
	}

	time.Sleep(10 * time.Second)

}
