package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/wunicorns/goutils/batch"
	"github.com/wunicorns/goutils/metrics"
)

var (
	measureTimeDesc = prometheus.NewDesc(
		"task_terminate_measure_time",
		"task_terminate_measure_time",
		[]string{"jobname"},
		nil)
)

func main() {

	// batch.Module()
	// metrics.MetricServer()

	// chanReceiver{}

	// receive := make(chan interface{})

	// go func() {
	// 	receive <- 10.
	// }()

	// num := <-receive
	// fmt.Println(num)
	// num2 := num.(float64) * 0.
	// fmt.Println(0 / num2)

	// fmt.Println("done")

	promConfig := metrics.PromConfig{
		MetricsPath:        "/metrics",
		ListenAddress:      ":9999",
		WebListenAddresses: []string{":9999"},
		WebSystemdSocket:   false,
		WebConfigFile:      "",
	}

	var sourceDesc = metrics.ScrapeSourceDesc{
		"measure_time": measureTimeDesc,
	}

	data := make(map[string]float64)

	data["measure_time"] = 1234

	launcher := metrics.MetricLauncher{
		Data:       &data,
		MetricDesc: &sourceDesc,
	}

	launcher.Serve(&promConfig)

}

type Task interface {
	Run()
}

type FuncTask func()

func (f FuncTask) Run() { f() }

func Call(cmd Task) {
	cmd.Run()
}

func _main() {
	callable := func() {
		fmt.Println(1)
	}
	Call(FuncTask(callable))
}
