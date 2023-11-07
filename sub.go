package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	fmt.Println("is sub routune. ", time.Now())
	log.Println("sub task")
}

func CallMe() {
	log.Println("test !!")
}
