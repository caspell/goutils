package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	promVersion "github.com/prometheus/common/version"
)

type (
	scrapeRequest struct {
		results chan<- prometheus.Metric
		done    chan struct{}
	}

	Source struct {
		Name  string
		Value int
	}

	ScrapeSourceDesc map[string]*prometheus.Desc
	ScrapeSource     func()

	MetricCollector struct {
		scrapeChan chan scrapeRequest
		Source     func() *map[string]float64
		SourceDesc ScrapeSourceDesc
	}
)

var version string

func init() {
	promVersion.Version = version
	prometheus.MustRegister(promVersion.NewCollector("application_exporter"))
}

func NewProcessCollector(source func() *map[string]float64, sourceDesc ScrapeSourceDesc) (*MetricCollector, error) {

	p := &MetricCollector{
		scrapeChan: make(chan scrapeRequest),
		Source:     source,
		SourceDesc: sourceDesc,
	}

	go p.start()

	return p, nil
}

func (p *MetricCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, desc := range p.SourceDesc {
		ch <- desc
	}
}

func (p *MetricCollector) Collect(ch chan<- prometheus.Metric) {
	req := scrapeRequest{results: ch, done: make(chan struct{})}
	p.scrapeChan <- req
	<-req.done
}

func (p *MetricCollector) start() {
	for req := range p.scrapeChan {
		ch := req.results
		p.scrape(ch)
		req.done <- struct{}{}
	}
}

// func (p *MetricCollector) Update () error {

// 	p.S
// }

func (p *MetricCollector) scrape(ch chan<- prometheus.Metric) {

	// if err := p.Update(); err != nil {
	// 	log.Printf("scrape error : %v", err)
	// }

	sources := *p.Source()
	sourcesDesc := p.SourceDesc

	for gname, source := range sources {
		ch <- prometheus.MustNewConstMetric(sourcesDesc[gname], prometheus.GaugeValue, float64(source), gname)
	}
}
