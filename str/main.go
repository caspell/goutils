package main

import (
	"fmt"
	"strings"
)

func Split(text, sep string) (string, string) {
	key, val, _ := strings.Cut(text, sep)
	return strings.TrimSpace(key), strings.TrimSpace(val)
}

func init() {
	fmt.Println("init")
}

func main() {

	value := "test 4444"

	vs := strings.Split(value, " ")

	fmt.Println(vs[len(vs)-1])
	// values1, values2 := Split(value, " ")
	// strings.Split()
	// strings.Index(value, "No")

	fmt.Println(len(vs))

	// fmt.Printf("%s\n%s", values1, values2)
}
