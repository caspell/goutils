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

func main() {

	obj := &ApiResult{}

	val := `{"result":"ONGOING"a}`

	if err := json.Unmarshal([]byte(val), obj); err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("%v\n", obj.Result)

	fmt.Println("done")

	fmt.Println(obj == nil)
}
