package log4libwrapper

import (
	"github.com/anthonyraymond/go-log4lib"
	"github.com/sirupsen/logrus"
)

type logrusLoggerWrapper struct {
	logrus.Logger
}

// Don't explicitly override method because logrus does not support caller depth => https://github.com/sirupsen/logrus/issues/972
func WrapLogrusLogger(pointerToLogger *logrus.Logger) log4lib.LibLogger {
	return pointerToLogger
}
