package config

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func startWatching() {
	viper.WatchConfig()
	viper.OnConfigChange(onConfigChangeEvent)
}

func onConfigChangeEvent(event fsnotify.Event) {
	log.Debug("Config change detected: \"", event.Name, "\"")
	LoadAll()
}
