package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	measureTimeDesc = prometheus.NewDesc(
		"task_terminate_measure_time",
		"task_terminate_measure_time",
		[]string{"jobname"},
		nil)
)

func TestServe(t *testing.T) {

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
