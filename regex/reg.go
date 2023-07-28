package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	processName := "ssh"

	r, _ := regexp.Compile(`[ {()}~!@#$%^&*/|'"]+`)

	pname := r.ReplaceAllString(processName, "")

	fmt.Println("done ", pname)

	var cmdOption string
	if !strings.EqualFold(pname, "") {
		cmdOption = strings.Join(([]string{" | grep ", pname}), "")
	}

	fmt.Println(cmdOption)

	// fmt.Println(fmt.Sprintf(_PROCESS_CMD, "agent"))
}
