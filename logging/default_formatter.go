package logging

import (
	"fmt"
	"time"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

type DefaultFormatter struct {
}

func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

func (f *DefaultFormatter) FormatMessage(exctx exctx.ExecutionContext, level LogLevel, msg string) string {
	strTime := currentTimeString()
	sLevel := logLevelString(level)

	return fmt.Sprintf("%s (%s) %s - %s", strTime, exctx.ID.String()[:8], sLevel, msg)
}

func logLevelString(level LogLevel) string {
	var sLevel string

	switch level {
	case Debug:
		sLevel = "D"
	case Log:
		sLevel = "L"
	case Warn:
		sLevel = "W"
	case Error:
		sLevel = "E"
	case Fatal:
		sLevel = "F"
	}

	return sLevel
}

func currentTimeString() string {
	loc, _ := time.LoadLocation("UTC")
	t := time.Now().In(loc)

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d.%03dZ",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/100000)
}
