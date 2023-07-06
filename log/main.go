package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const MAX_TRACE = 10

func main() {

	fmt.Println(getFullCaller(nil))

}

func getFullCaller(aaa interface{}) string {

	a, b := aaa.(string)
	fmt.Println(a, b)

	return ""
}

func getCallerTrace(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".")
	}
	return fmt.Sprintf("[%s:%d %s]", filepath.Base(file), line, fnName)
}

func test1() {
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log := logrus.New()
	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}
