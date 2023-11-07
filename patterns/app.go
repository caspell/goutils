package patterns

import (
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	_ "golang.org/x/sys/unix"
	"golang.org/x/sys/windows"
)

func Call1() {

	num1, num2 := 0, 0
	goto loop1

loop1:
	for {
		num1++
		fmt.Printf("%d / %d \n", windows.GetCurrentProcessId(), windows.GetCurrentThreadId())
	loop2:
		for {
			num2++
			fmt.Println(num2)
			if num2 == 5 {
				break loop2
			}
		}
		fmt.Println("after Loop2")
		break loop1
	}

	fmt.Println("after Loop1")
}

func Call2(v string) {
	if v == "m1" {
		goto m1
	} else if v == "m2" {
		goto m2
	} else if v == "m3" {
		goto m3
	}

m1:
	fmt.Println("m1 called!")
	goto end
m2:
	fmt.Println("m2 called!")
	goto end
m3:
	fmt.Println("m3 called!")
	goto end
end:
}

func dateCheck() {

	tsub1 := time.Now()

	tsub := time.Now().Add(time.Second * time.Duration(300))

	dur := tsub.Sub(time.Now().Add(time.Second * time.Duration(301)))

	fmt.Println(fmt.Sprintf("%d-%d-%d %d:%d:%d", tsub1.Year(), tsub1.Month(), tsub1.Day(), tsub1.Hour(), tsub1.Minute(), tsub1.Second()))
	fmt.Println(fmt.Sprintf("%d-%d-%d %d:%d:%d", tsub.Year(), tsub.Month(), tsub.Day(), tsub.Hour(), tsub.Minute(), tsub.Second()))

	fmt.Println(dur > 0)

}

func mapCheck() {

	defer func() {
		v := recover()
		fmt.Printf("recoevered: %v \n", v)
		values := strings.Split(string(debug.Stack()), "\n")
		var panicIndex int = 0
		var line string
		for panicIndex, line = range values {
			if strings.HasPrefix(line, "panic") {
				panicIndex++
				break
			}
		}
		fmt.Printf("%v", values[panicIndex+1:])
	}()

	map1 := make(map[string]interface{})

	map1["key1"] = "11234"
	map1["key2"] = "21234"
	map1["key3"] = "31234"

	if v, exist := map1["key1"].(string); exist {
		fmt.Println(fmt.Sprintf("exist :: %s, %v", v, exist))
	} else {
		fmt.Println(fmt.Sprintf("not exist ::%s, %v", v, exist))
	}

	if v, exist := map1["key4"].(string); exist {
		fmt.Println(fmt.Sprintf("exist :: %s, %v", v, exist))
	} else {
		fmt.Println(fmt.Sprintf("not exist ::%s, %v", v, exist))
	}

	value := map1["key5"].(string)

	fmt.Println(value)

	fmt.Println("done!")

}

type SubConfig struct {
	Value string `toml:"VALUE"`
}

type Config struct {
	Cfgs []SubConfig
}

func DateFormat(date string) string {
	_, s := time.Now().Zone() //time.Now().Zone()
	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Println(err)
		return "0001-01-01 00:00:00"
	}
	parsed, err := time.ParseInLocation(time.RFC3339, date, loc)
	if err != nil {
		fmt.Println(err)
		return "0001-01-01 00:00:00"
	}
	duration, err := time.ParseDuration(fmt.Sprintf("%ds", s))
	if err != nil {
		fmt.Println(err)
		return "0001-01-01 00:00:00"
	}
	parsed = parsed.Add(duration)
	return parsed.Format("2006-01-02 15:04:05")
}

func timeTest() {
	date := "2023-02-17T01:38:03.294Z"

	formatedDt := DateFormat(date)

	fmt.Println(formatedDt)

	// format := "2006-01-02 15:04:05"
	// tz, s := time.Now().Zone()
	// loc, _ := time.LoadLocation(tz)
	// parsed, _ := time.ParseInLocation(time.RFC3339, date, loc)
	// duration, _ := time.ParseDuration(fmt.Sprintf("%ds", s))
	// parsed = parsed.Add(duration)
	// fmt.Println(parsed.Format(format))
}

func callTest(param1 string, params2 ...string) {

	fmt.Println(param1)
	// fmt.Println(params2)
	for i, v := range params2 {
		fmt.Printf(" %d, %s \n", i, v)
	}
}

type SerialData struct {
	Measurement string
	Value       float64
	Timestamp   int64
}

type SerialDatas []*SerialData

const (
	MetricNameLabel  = "__name__"
	MetricValueLabel = "_value"
	MetricTimeLabel  = "_time"
	ReplaceDelimeter = "=>"
)

const (
	minimumTick  = time.Millisecond
	second       = int64(time.Second / minimumTick)
	nanosPerTick = int64(minimumTick / time.Nanosecond)
)

func (s *SerialData) Time() time.Time {
	return time.Unix(int64(s.Timestamp)/second, (int64(s.Timestamp)%second)*nanosPerTick)
}

func Main() {
	// var configFilePath string

	// flag.StringVar(&configFilePath, "c", "config.toml", "(*Require) Config file")

	// flag.Parse()

	// var appConfig Config

	// if _, err := toml.DecodeFile(configFilePath, &appConfig); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%v \n", strings.Split("test.toml,123.toml", ","))

	// testMap := make(map[string]interface{})

	// testMap["key1"] = 123
	// testMap["key2"] = "testok!"
	// testMap["key3"] = 123.

	// for k, v := range testMap {
	// 	fmt.Println(k, v)

	// 	aaa := reflect.TypeOf(v)

	// 	switch aaa.Kind() {
	// 	case reflect.String:
	// 		fmt.Println("is string")
	// 	case reflect.Int64:
	// 		fmt.Println("is number")
	// 	case reflect.Float64:
	// 		fmt.Println("is float")
	// 	default:
	// 		fmt.Println("default")
	// 	}

	// }

	// params := []string{"test1"}

	// params = append(params, "test2")

	// callTest("first", "value1", "value2")
	// fmt.Println("-----------")
	// callTest("second", params...)

	// sd := &SerialData{
	// 	Timestamp: 1676529408899,
	// }

	// fmt.Println(sd.Time())

	timeTest()
}
