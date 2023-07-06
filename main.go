package main

import (
	"fmt"

	_ "github.com/wunicorns/goutils/batch"
	"github.com/wunicorns/goutils/httprequest"
	_ "github.com/wunicorns/goutils/querybuilder"
)

func main() {

	// fmt.Println("start")

	// filename := "querybuilder/sample.yml"

	// querybuilder.QueryBuilderInitialize(filename)

	// stmt := querybuilder.GetQueryBuilder().GetStatement("get_sample_list")

	// fmt.Println(stmt.Script)

	fmt.Println("start")

	if str, err := httprequest.GetJson("", nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str)
	}
}
