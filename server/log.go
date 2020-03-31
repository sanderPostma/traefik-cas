package server

import (
	"github.com/sirupsen/logrus"
	golog "log"
	"os"
)

var log *logrus.Logger

func NewDefaultLogger() *logrus.Logger {
	// Setup logger
	log = logrus.StandardLogger()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	golog.SetOutput(log.Writer())
	return log
}
