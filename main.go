package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/seanarwa/common/config"
	"github.com/seanarwa/ds/api"
	"github.com/seanarwa/ds/db/mongo"
	log "github.com/sirupsen/logrus"
)

func main() {

	helpFlag, logConfigFile, configFile := config.ParseArgs()

	if helpFlag {
		config.PrintUsage()
		os.Exit(0)
	}

	config.Set("cmd.log_config_file", logConfigFile)
	config.Set("cmd.config_file", configFile)
	config.Init()

	log.Info(config.GetString("name"), " v", config.GetString("version"), " has started")

	mongo.Connect(config.GetString("db.mongo.url"))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		api.Start()
	}()

	<-done
	mongo.Disconnect()
	log.Info(config.GetString("name"), " v", config.GetString("version"), " has stopped")
}
