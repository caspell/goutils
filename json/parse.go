package main

import (
	"encoding/json"
	"fmt"
)

type ApiResult struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func main2() {

	obj := &ApiResult{}
	val := `{"result":"ONGOING"a}`
	if err := json.Unmarshal([]byte(val), obj); err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("%v\n", obj.Result)
	fmt.Println("done")
	fmt.Println(obj == nil)

}

func main() {

	map1 := make(map[string]int)

	map1["b111"] = 222
	map1["a111"] = 111
	map1["d111"] = 333
	map1["c111"] = 444

	for k, v := range map1 {
		fmt.Println(k, v)
	}

}
