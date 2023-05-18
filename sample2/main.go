package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
)

func protoToSamples(req *prompb.WriteRequest) {
	fmt.Printf("%v", req)

	// var samples model.Samples
	// for _, ts := range req.Timeseries {
	// 	fmt.Println(ts.Labels)
	// 	fmt.Println(ts.Samples)
	// 	// metric := make(model.Metric, len(ts.Labels))
	// 	// for _, l := range ts.Labels {
	// 	// 	metric[model.LabelName(l.Name)] = model.LabelValue(l.Value)
	// 	// }
	// 	// for _, s := range ts.Samples {
	// 	// 	samples = append(samples, &model.Sample{
	// 	// 		Metric:    metric,
	// 	// 		Value:     model.SampleValue(s.Value),
	// 	// 		Timestamp: model.Time(s.Timestamp),
	// 	// 	})
	// 	// }
	// }
}

func main() {

	fmt.Println("run")

	logFilePath := "D:\\windows_exporter_process.log"

	if file, err := os.ReadFile(logFilePath); err != nil {
		panic(err)
	} else {

		lines := strings.Split(string(file), "\n")
		// fmt.Println(string(file))
		_map := make(map[string]int)

		//reg, _ := regexp.Compile("^[_a-z]+([ ]|[{])")

		for _, l := range lines {
			// fmt.Println(l)

			// if !strings.HasPrefix(l, "#") {
			// 	b := reg.Find([]byte(l))
			// 	_map[string(b[:len(b)-1])] += 1
			// }

			if strings.HasPrefix(l, "# HELP") {
				fmt.Println(l)
			}
		}

		for k, v := range _map {
			fmt.Println(k, v)
		}

	}

}
