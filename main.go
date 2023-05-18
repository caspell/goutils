package main

import (
	"fmt"

	"github.com/wunicorns/goutils/batch"
	"github.com/wunicorns/goutils/metrics"
)

func main() {

	fmt.Println("run")

	batch.Module()

	metrics.MetricServer()

}
