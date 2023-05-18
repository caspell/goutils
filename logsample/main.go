package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
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
