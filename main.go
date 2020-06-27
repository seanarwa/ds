package main

import (
	"os"

	"github.com/seanarwa/ds/config"
	"github.com/seanarwa/ds/db/mongo"
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

	mongo.Connect(config.GetString("db.mongo.url"))
	mongo.Test()
	mongo.Disconnect()
}
