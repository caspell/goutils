package main

import (
	"fmt"
	"sync"
	// 	_ "github.com/wunicorns/goutils/batch"
	// 	_ "github.com/wunicorns/goutils/httprequest"
	// 	_ "github.com/wunicorns/goutils/patterns"
	// 	_ "github.com/wunicorns/goutils/querybuilder"
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

	// for i, _ := range make([]int, 10) {
	// 	go patterns.Execute(fmt.Sprintf("id: %d", i))
	// }

	// time.Sleep(10 * time.Second)

	// fmt.Println(int64(math.Pow(2, 10)))

	// fmt.Println(8192 * int64(math.Pow(2, 10)))

	// fmt.Println(8 * 1024 * 1024)

	// unit := strings.ToUpper("gb")
	// basisUnits := []string{
	// 	"Bytes", "KB", "MB", "GB", "TB", "PB",
	// }

	// fmt.Println(IndexOf(basisUnits, unit))

	wg := &sync.WaitGroup{}

	mu := &sync.Mutex{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		mu.Lock()
		go func(index int) {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()
			fmt.Println(index)
		}(i)
	}

	wg.Wait()

	fmt.Println("done!")
}
