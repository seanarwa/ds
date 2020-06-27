package config

import (
	"encoding/json"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

func parseFilepath(filepath string) (ext string, filename string, dir string) {
	logConfigFile := path.Clean(filepath)
	ext = path.Ext(logConfigFile)
	ext = strings.Replace(ext, ".", "", 1)
	base := path.Base(logConfigFile)
	idx := strings.LastIndex(base, ".")
	filename = base
	if idx != -1 {
		filename = base[0:idx]
	}
	dir = path.Dir(logConfigFile)

	return ext, filename, dir
}

func prettyPrint(config map[string]interface{}) {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal("Error occured when trying to pretty print config: ", err)
	}
	log.Trace(" " + string(b))
}
