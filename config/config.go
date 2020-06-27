package config

import (
	"github.com/seanarwa/ds/logging"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Init() {
	setAllDefaults()

	LoadAll()

	startWatching()
	log.Debug("Configuration change watch has started")
}

func LoadAll() {
	logConfigFile := GetString("cmd.log_config_file")
	loadLogging(logConfigFile)

	configFile := GetString("cmd.config_file")
	loadMain(configFile)

	log.WithFields(log.Fields{
		"log_config_file": logConfigFile,
		"config_file":     configFile,
	}).Debug("All configuration loaded")
	prettyPrint(viper.AllSettings())
}

func loadMain(configFile string) {
	err := addConfig(configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"config_file": configFile,
		}).Fatal("Error occured when trying to read in loaded config: ", err)
	}
}

func loadLogging(configFile string) {
	err := addConfig(configFile)
	if err != nil {
		log.WithFields(log.Fields{
			"log_config_file": configFile,
		}).Fatal("Error occured when trying to read in loaded logging config: ", err)
	}

	logging.Init(
		GetString("logging.file"),
		GetString("logging.level"),
		GetBool("logging.append"),
	)
}

func Get(key string) interface{} {
	safeCheck(key)
	return viper.Get(key)
}

func GetString(key string) string {
	safeCheck(key)
	return viper.GetString(key)
}

func GetBool(key string) bool {
	safeCheck(key)
	return viper.GetBool(key)
}

func GetInt(key string) int {
	safeCheck(key)
	return viper.GetInt(key)
}

func GetFloat64(key string) float64 {
	safeCheck(key)
	return viper.GetFloat64(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}

func Exists(key string) bool {
	return viper.IsSet(key)
}

func safeCheck(key string) {
	if !Exists(key) {
		log.WithFields(log.Fields{
			"key": key,
		}).Fatal("Error occured when trying to get key from config: key does not exists")
	} else {
		log.WithFields(log.Fields{
			"key": key,
		}).Trace("Config key safe check ok")
	}
}

func addConfig(filepath string) (err error) {
	ext, filename, dir := parseFilepath(filepath)
	viper.AddConfigPath(dir)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext)
	err = viper.MergeInConfig()
	if err != nil {
		log.Error("Error occured when trying to add file to config: ", err)
	} else {
		log.WithFields(log.Fields{
			"file": filepath,
		}).Trace("Added config")
	}
	return err
}
