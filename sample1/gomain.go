package main

import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"

	_ "github.com/wunicorns/goutils/patterns"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func _main() {

	val := `{}`

	obj := Data{
		Name: "myname",
		Age:  10,
	}

	if bytes, err := json.Marshal(obj); err != nil {
		panic(err)
	} else {
		fmt.Println("marshal")
		fmt.Println(string(bytes))
	}

	bval := []byte(val)

	obj2 := Data{}

	json.Unmarshal(bval, &obj2)

	fmt.Println("done")

	// patterns.PatternSingleTone()

}

func FuncTest(name string) func() {
	num := rand.Int()
	fmt.Println("intask 1 ", name, num)
	begin := time.Now()
	return func() {
		time.Sleep(1 * time.Second)
		last := time.Since(begin).Milliseconds()
		fmt.Println("intask 2 ", name, last, num)
		err := recover()
		if err != nil {
			fmt.Println("error :: ", err)
		}
	}
}

func main() {

	fmt.Println("task 1")

	defer FuncTest("test1")()
	defer FuncTest("test2")()
	defer FuncTest("test3")()
	defer FuncTest("test4")()
	// if true {
	// 	panic("func")
	// }
	fmt.Println("task 2")
}
