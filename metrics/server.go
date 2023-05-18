package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricServer() {

	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Create new metrics and register them using the custom registry.
	m := NewMetrics(reg)

	// Set values for the new created metrics.
	m.CpuTemp.Set(65.3)
	m.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	log.Fatal(http.ListenAndServe(":8080", nil))

}
