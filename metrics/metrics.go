package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	CpuTemp    prometheus.Gauge
	HdFailures *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		CpuTemp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU.",
		}),
		HdFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hd_errors_total",
				Help: "Number of hard-disk errors.",
			},
			[]string{"device"},
		),
	}
	reg.MustRegister(m.CpuTemp)
	reg.MustRegister(m.HdFailures)
	return m
}
