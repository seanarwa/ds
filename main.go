package main

import (
	"os"

	"seanchang.me/ds/config"
	"seanchang.me/ds/db/mongo"
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
