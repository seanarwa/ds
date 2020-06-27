package config

import (
	"github.com/spf13/viper"
)

var defaults = map[string]interface{}{
	"logging": map[string]interface{}{
		"append": true,
		"file":   "log/main.log",
		"level":  "debug",
	},
	"cmd": map[string]interface{}{
		"log_config_file": "conf/log.yaml",
		"config_file":     "conf/app.yaml",
	},
}

func setAllDefaults() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}
