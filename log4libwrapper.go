package log4libwrapper

import (
	"github.com/anthonyraymond/go-log4lib"
	"github.com/sirupsen/logrus"
)

type logrusLoggerWrapper struct {
	delegate *logrus.Logger
}

func (w *logrusLoggerWrapper) Info(args ...interface{}) {
	w.delegate.Info(args)
}

func (w *logrusLoggerWrapper) Warn(args ...interface{}) {
	w.delegate.Warn(args)
}

func (w *logrusLoggerWrapper) Error(args ...interface{}) {
	w.delegate.Error(args)
}

func (w *logrusLoggerWrapper) Panic(args ...interface{}) {
	w.delegate.Panic(args)
}

func (w *logrusLoggerWrapper) Fatal(args ...interface{}) {
	w.delegate.Fatal(args)
}

func (w *logrusLoggerWrapper) Infof(template string, args ...interface{}) {
	w.delegate.Infof(template, args)
}

func (w *logrusLoggerWrapper) Warnf(template string, args ...interface{}) {
	w.delegate.Warnf(template, args)
}

func (w *logrusLoggerWrapper) Errorf(template string, args ...interface{}) {
	w.delegate.Errorf(template, args)
}

func (w *logrusLoggerWrapper) Panicf(template string, args ...interface{}) {
	w.delegate.Panicf(template, args)
}

func (w *logrusLoggerWrapper) Fatalf(template string, args ...interface{}) {
	w.delegate.Fatalf(template, args)
}

func WrapLogrusLogger(pointerToLogger *logrus.Logger) log4lib.LibLogger {
	return &logrusLoggerWrapper{delegate: pointerToLogger}
}
