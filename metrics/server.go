// package metrics

// import (
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// func MetricServer() {

// 	// Create a non-global registry.
// 	reg := prometheus.NewRegistry()

// 	// Create new metrics and register them using the custom registry.
// 	m := NewMetrics(reg)

// 	// Set values for the new created metrics.
// 	go func() {
// 		for {
// 			m.CpuTemp.Set(rand.Float64())
// 			m.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// 	// Expose metrics and custom registry via an HTTP server
// 	// using the HandleFor function. "/metrics" is the usual endpoint for that.
// 	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }
