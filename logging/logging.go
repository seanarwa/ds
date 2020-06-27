package logging

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

var initialized = false
var allowedLevels = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

func Init(file string, level string, append bool) {

	logrus.SetFormatter(&Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% %caller% [%level%]: %message% %fields%",
	})
	logrus.SetLevel(toLogrusLevel(level))
	logrus.SetReportCaller(true)

	f := openLogFile(file, append)
	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)

	logrus.WithFields(logrus.Fields{
		"level": level,
		"file":  file,
	}).Debug("Initialized logger successfully")
}

func openLogFile(path string, append bool) *os.File {

	dir, filename := filepath.Split(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}

	mode := os.O_RDWR | os.O_CREATE
	if append {
		mode = mode | os.O_APPEND
	}
	f, err := os.OpenFile(path, mode, 0777)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"directory": dir,
			"filename":  filename,
		})
		logrus.Fatalf("Error occured when trying to open log file: %v", err)
	}

	return f
}

func toLogrusLevel(level string) logrus.Level {

	level = strings.ToLower(level)

	for k, v := range allowedLevels {
		if k == level {
			return v
		}
	}

	logrus.WithFields(logrus.Fields{
		"level_received":  level,
		"allowed_options": allowedLevels,
	}).Fatal("Error converting to logrus level")

	return logrus.FatalLevel
}
