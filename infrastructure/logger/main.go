package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func init() {
	level := strings.ToLower(os.Getenv("LOG_LEVEL"))

	switch level {
	case log.PanicLevel.String():
		log.SetLevel(log.PanicLevel)
	case log.FatalLevel.String():
		log.SetLevel(log.FatalLevel)
	case log.ErrorLevel.String():
		log.SetLevel(log.ErrorLevel)
	case log.WarnLevel.String():
		log.SetLevel(log.WarnLevel)
	case log.InfoLevel.String():
		log.SetLevel(log.InfoLevel)
	case log.DebugLevel.String():
		log.SetLevel(log.DebugLevel)
	case log.TraceLevel.String():
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
