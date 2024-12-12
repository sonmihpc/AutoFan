package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mode       string      `mapstructure:"mode" yaml:"mode"`
	Interval   int         `mapstructure:"interval" yaml:"interval"`
	Thresholds []Threshold `mapstructure:"thresholds" yaml:"thresholds"`
	PwdIds     []int       `mapstructure:"pwd-ids" yaml:"pwd-ids"`
	GpuDebug   bool        `mapstructure:"gpu-debug" yaml:"gpu-debug"`
	IpmiDebug  bool        `mapstructure:"ipmi-debug" yaml:"ipmi-debug"`
}

func readConf(path ...string) Config {
	var config Config
	var configFile string
	if len(path) == 0 {
		flag.StringVar(&configFile, "c", "", "Read config from the specified file.")
		flag.Parse()
		if configFile == "" {
			panic("please set the config file by -c [file].")
		} else {
			fmt.Printf("use the config file from command flag.\n")
		}
	} else {
		configFile = path[0]
		fmt.Printf("use the configFile specified by argument.\n")
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic("Fail to read config file.")
	}

	if err := v.Unmarshal(&config); err != nil {
		panic("Fail to unmarshal config.")
	}

	return config
}
