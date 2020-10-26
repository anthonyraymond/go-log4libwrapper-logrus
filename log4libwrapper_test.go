package log4libwrapper

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"testing"
)

func TestWrapLogrusLogger_MustLog(t *testing.T) {
	logReceiver := &strings.Builder{}

	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.DebugLevel)
	logrusLogger.SetOutput(io.MultiWriter(logReceiver, os.Stdout))
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:          true,
		DisableSorting:            true,
		ForceQuote: true,
	})

	logger := WrapLogrusLogger(logrusLogger)
	logger.Debug("coucou")
	logger.Info("coucou")
	logger.Warn("coucou")
	logger.Error("coucou")
	logger.Debug("coucou", "joe", "la bidouille")
	logger.Info("coucou", "joe", "la bidouille")
	logger.Warn("coucou", "joe", "la bidouille")
	logger.Error("coucou", "joe", "la bidouille")
	logger.Debugf("coucou %s :)", "joe")
	logger.Infof("coucou %s :)", "joe")
	logger.Warnf("coucou %s :)", "joe")
	logger.Errorf("coucou %s :)", "joe")
	logger.Debugf("coucou %s %d :)", "joe", 12)
	logger.Infof("coucou %s %d :)", "joe", 12)
	logger.Warnf("coucou %s %d :)", "joe", 12)
	logger.Errorf("coucou %s %d :)", "joe", 12)

	expected := []string{
		"level=\"debug\" msg=\"coucou\"",
		"level=\"info\" msg=\"coucou\"",
		"level=\"warning\" msg=\"coucou\"",
		"level=\"error\" msg=\"coucou\"",
		"level=\"debug\" msg=\"coucoujoela bidouille\"",
		"level=\"info\" msg=\"coucoujoela bidouille\"",
		"level=\"warning\" msg=\"coucoujoela bidouille\"",
		"level=\"error\" msg=\"coucoujoela bidouille\"",
		"level=\"debug\" msg=\"coucou joe :)\"",
		"level=\"info\" msg=\"coucou joe :)\"",
		"level=\"warning\" msg=\"coucou joe :)\"",
		"level=\"error\" msg=\"coucou joe :)\"",
		"level=\"debug\" msg=\"coucou joe 12 :)\"",
		"level=\"info\" msg=\"coucou joe 12 :)\"",
		"level=\"warning\" msg=\"coucou joe 12 :)\"",
		"level=\"error\" msg=\"coucou joe 12 :)\"",
		"",
	}

	if !strings.EqualFold(strings.Join(expected, "\n"), logReceiver.String()) {
		t.Fatalf("log output is not correct, expected:\n%s\nactual:\n%s", strings.Join(expected, "\n"), logReceiver.String())
	}
}

func TestWrapZapLogger_CallerMustBeUnwrapped(t *testing.T) {
	// when the caller is added to the zap output we don't want to see the Info, Warn, ... wrapping method from zapLoggerWrapper, we want the real caller
	logReceiver := &strings.Builder{}


	logrusLogger := logrus.New()
	logrusLogger.SetReportCaller(true)
	logrusLogger.SetOutput(io.MultiWriter(logReceiver, os.Stdout))
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:          true,
		DisableSorting:            true,
	})

	WrapLogrusLogger(logrusLogger).Info("coucou")

	if !strings.Contains(logReceiver.String(), "go-log4libwrapper-logrus/log4libwrapper_test.go") {
		t.Fatal("caller is now correct, the wrapper should be ignored and the caller must be the real calling function (here go-log4libwrapper-zap/log4libwrapper_test.go)")
	}
}
