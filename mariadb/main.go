package mariadb

import "fmt"

func Run() {

	fmt.Println("run")

	list, _ := SelectServerList("")

	for _, r := range *list {
		fmt.Println(r)
	}

	fmt.Println("done")
}
