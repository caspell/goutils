package metrics

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/exporter-toolkit/web"
	_ "github.com/wunicorns/goutils/batch"
)

type PromConfig struct {
	MetricsPath        string
	ListenAddress      string
	WebListenAddresses []string
	WebSystemdSocket   bool
	WebConfigFile      string
}

type MetricLauncher struct {
	Data       *map[string]float64
	MetricDesc *ScrapeSourceDesc
}

func (m MetricLauncher) Serve(config *PromConfig) {

	var source = func() *map[string]float64 {
		return m.Data
	}

	if collector, err := NewProcessCollector(source, *m.MetricDesc); err != nil {
		panic(err)
	} else {
		prometheus.MustRegister(collector)
	}

	var tlsConfigFile web.FlagConfig = web.FlagConfig{
		WebListenAddresses: &config.WebListenAddresses,
		WebSystemdSocket:   &config.WebSystemdSocket,
		WebConfigFile:      &config.WebConfigFile,
	}

	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	http.Handle(config.MetricsPath, promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>TTM batch Exporter</title></head>
			<body>
			<h1>TTM batch Exporter</h1>
			<p><a href="` + config.MetricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	server := &http.Server{
		Addr: config.ListenAddress,
	}
	if err := web.ListenAndServe(server, &tlsConfigFile, logger); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
		os.Exit(1)
	}

}
