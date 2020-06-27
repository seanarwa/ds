package logging

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "%time% %caller% [%level%]: %message%%fields%"
	defaultTimestampFormat = time.RFC3339
)

var spacing = 0

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%level%", level, 1)

	caller := entry.Caller.File
	caller = caller[strings.LastIndex(caller, "/")+1 : len(caller)]
	if len(caller) > spacing {
		spacing = len(caller)
	}
	caller += strings.Repeat(" ", spacing-len(caller))
	output = strings.Replace(output, "%caller%", caller, 1)

	header := strings.Replace(output, "%message%", "", 1)
	header = strings.Replace(header, "%fields%", "", 1)

	output = strings.Replace(output, "%message%", entry.Message, 1)

	fields := ""
	for k, val := range entry.Data {

		switch v := val.(type) {
		case string:
			fields += fmt.Sprintf(`%s="%s" `, k, v)
		case int:
			s := strconv.Itoa(v)
			fields += fmt.Sprintf(`%s=%s `, k, s)
		case bool:
			s := strconv.FormatBool(v)
			fields += fmt.Sprintf(`%s=%s `, k, s)
		}
	}
	if len(fields) > 0 {
		fields = fields[:len(fields)-1]
	}
	output = strings.Replace(output, "%fields%", fields, 1)

	output = strings.ReplaceAll(output, "\n", "\n"+header)
	output += "\n"

	return []byte(output), nil
}
