module github.com/wunicorns/goutils

replace (
	github.com/wunicorns/goutils/batch => ./batch
	github.com/wunicorns/goutils/mail => ./mail
	github.com/wunicorns/goutils/metrics => ./metrics
	github.com/wunicorns/goutils/patterns => ./patterns
)

go 1.19

require (
	github.com/prometheus/common v0.44.0
	github.com/prometheus/prometheus v0.45.0
	github.com/sirupsen/logrus v1.9.3
	github.com/wunicorns/goutils/batch v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.9.0
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
)
