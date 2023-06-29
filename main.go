package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/wunicorns/goutils/batch"
)

func main() {

	fmt.Println("start")

	filename := "querybuilder/sample.yml"

	if finfo, err := os.Stat(filename); err != nil {
		panic(err)
	} else {

		fmt.Println(finfo.Name(), strings.HasSuffix(finfo.Name(), ".yml"))
	}

}
