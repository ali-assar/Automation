package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func Init(logLevel, logFormat string) {
	log = logrus.New()
	log.Out = os.Stdout

	level, err := logrus.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		level = logrus.InfoLevel
	}
	log.SetLevel(level)

	if logFormat == "json" {
		log.SetFormatter(&logrus.JSONFormatter{})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}
}

func Info(args ...any)  { log.Info(args...) }
func Warn(args ...any)  { log.Warn(args...) }
func Error(args ...any) { log.Error(args...) }
func Debug(args ...any) { log.Debug(args...) }

// WithField / WithFields if you want structured logging
func WithField(key string, value any) *logrus.Entry {
	return log.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}
