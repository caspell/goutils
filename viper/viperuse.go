package viper

import (
	"embed"
	"flag"
	"fmt"
	_ "runtime/debug"
	_ "sync"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/wunicorns/goutils/api"
	// "github.com/wunicorns/goutils/crypto"
	// "github.com/wunicorns/goutils/mail"
	// "github.com/wunicorns/goutils/batch"
	// "github.com/wunicorns/goutils/channel"
	// "github.com/wunicorns/goutils/api"
)

type (
	Config struct {
		Metrics                 MetricsConfig
		Influxdb                InfluxdbConfig
		VTUNE_PROFILER_2022_DIR string
		vtune_profiler_2022_dir string
		vtuneProfiler2022Dir    string
	}
	InfluxdbConfig struct {
		Address  string `toml:"INFLUX_ADDRESS" env:"INFLUX_ADDRESS"`
		Token    string `toml:"INFLUX_TOKEN" env:"INFLUX_TOKEN"`
		Org      string `toml:"INFLUX_ORG" env:"INFLUX_ORG"`
		OrgId    string `toml:"INFLUX_ORG_ID" env:"INFLUX_ORG_ID"`
		User     string `toml:"INFLUX_USER" env:"INFLUX_USER"`
		Password string `toml:"INFLUX_PASSWORD" env:"INFLUX_PASSWORD"`
		Bucket   string `toml:"INFLUX_BUCKET" env:"INFLUX_BUCKET"`
	}
	MetricsConfig struct {
		MetricsPath        string
		ListenAddress      string
		WebListenAddresses []string
		WebSystemdSocket   bool
		WebConfigFile      string
	}
)

//go:embed files
var d embed.FS

func init() {
	fmt.Println("main package initialized. ", time.Now())

	viper.SetConfigType("toml")
	viper.SetConfigName("server")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// viper.RegisterAlias("influxdb.address", "INFLUX_ADDRESS")
	// viper.RegisterAlias("influxdb.token", "influxdb.influx_token")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	// if err := viper.BindEnv("INFLUX_ADDRESS"); err != nil {
	// 	panic(err)
	// }

	// INFLUX_ADDRESS
	// viper.RegisterAlias("INFLUX_ORG_ID", "influxdb.INFLUX_ORG_ID")

}

func main() {

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	flag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	simpleSns := &api.SimpleSNS{}

	simpleSns.Run()

}
