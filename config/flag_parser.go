package config

import (
	"os"

	"github.com/pborman/getopt/v2"
)

func ParseArgs() (help bool, log_config_file string, config_file string) {

	cmdDefaults := defaults["cmd"].(map[string]interface{})

	var help_opt = getopt.BoolLong("help", 'h', "Display usage")
	var log_config_file_opt = getopt.StringLong(
		"logging",
		'l',
		cmdDefaults["log_config_file"].(string),
		"The logging configuration YAML file",
	)
	var config_file_opt = getopt.StringLong(
		"config",
		'f',
		cmdDefaults["config_file"].(string),
		"The application configuration YAML file",
	)

	getopt.ParseV2()

	return *help_opt, *log_config_file_opt, *config_file_opt
}

func PrintUsage() {
	getopt.PrintUsage(os.Stdout)
}
