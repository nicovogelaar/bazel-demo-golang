package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func Error(fmt string, args ...interface{}) {
	log.Errorf(fmt, args...)
}

func Warning(fmt string, args ...interface{}) {
	log.Warningf(fmt, args...)
}

func Info(fmt string, args ...interface{}) {
	log.Infof(fmt, args...)
}

func Debug(fmt string, args ...interface{}) {
	log.Debugf(fmt, args...)
}
