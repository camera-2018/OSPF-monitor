package conf

import (
	"fmt"
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Metadata struct {
	Name string
}

var (
	Uptime = UptimeCfg{
		// Interval for uptime is statically set to
		Interval: time.Minute,
	}
	DSN            string
	Probes         []Probe
	Port           int
	Interval       time.Duration
	ProbeTimeout   time.Duration
	UpdateCallBack func()
)

func Init() error {
	viper.SetDefault("port", 8080)
	viper.SetDefault("interval", 60)
	viper.SetDefault("probeTimeout", 10)
	viper.SetDefault("uptime.recordDuration", "48h")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	viper.OnConfigChange(
		func(_ fsnotify.Event) {
			if err := update(); err != nil {
				fmt.Println("update probes fail", err)
			}
		})
	viper.WatchConfig()

	if err != nil {
		return err
	}
	Port = viper.GetInt("port")
	DSN = viper.GetString("dsn")

	return update()
}

func update() error {
	// update probe
	var tmp []Probe
	probes := viper.Get("probe").(map[string]any)
	for name, probe := range probes {
		probe, ok := probe.(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:%v", probe)
		}
		parser, ok := probe["parse"].(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid field parse")
		}
		fetcher, ok := probe["fetch"].(map[string]any)
		if !ok {
			return fmt.Errorf("parse config error:invalid field fetch")
		}
		tmp = append(tmp, Probe{
			Name:  name,
			Parse: parser,
			Fetch: fetcher,
		})
	}
	Probes = tmp

	Interval = time.Duration(viper.GetInt("interval")) * time.Second
	ProbeTimeout = time.Duration(viper.GetInt("probeTimeout")) * time.Second

	dur, err := time.ParseDuration(viper.GetString("uptime.store-duration"))
	if err != nil {
		return fmt.Errorf("parse uptime.store-duration fail:%v", err)
	}
	Uptime.StoreDuration = dur

	if UpdateCallBack != nil {
		UpdateCallBack()
	}
	log.Println("update config success")
	return nil
}
