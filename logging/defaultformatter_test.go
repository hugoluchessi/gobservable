package logging

import (
	"regexp"
	"testing"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

func TestDefaultFormatterDebug(t *testing.T) {
	exctx := exctx.Create()
	level := Debug
	msg := "Msg test"

	f := NewDefaultFormatter()

	fmtedMsg := f.FormatMessage(exctx, level, msg)
	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) D - Msg test"

	match, _ := regexp.MatchString(expectedRegex, fmtedMsg)

	if !match {
		t.Errorf("Test failed, '%s' has incorrect format", fmtedMsg)
	}
}

func TestDefaultFormatterLog(t *testing.T) {
	exctx := exctx.Create()
	level := Log
	msg := "Msg test"

	f := NewDefaultFormatter()

	fmtedMsg := f.FormatMessage(exctx, level, msg)
	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) L - Msg test"

	match, _ := regexp.MatchString(expectedRegex, fmtedMsg)

	if !match {
		t.Errorf("Test failed, '%s' has incorrect format", fmtedMsg)
	}
}

func TestDefaultFormatterWarn(t *testing.T) {
	exctx := exctx.Create()
	level := Warn
	msg := "Msg test"

	f := NewDefaultFormatter()

	fmtedMsg := f.FormatMessage(exctx, level, msg)
	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) W - Msg test"

	match, _ := regexp.MatchString(expectedRegex, fmtedMsg)

	if !match {
		t.Errorf("Test failed, '%s' has incorrect format", fmtedMsg)
	}
}

func TestDefaultFormatterError(t *testing.T) {
	exctx := exctx.Create()
	level := Error
	msg := "Msg test"

	f := NewDefaultFormatter()

	fmtedMsg := f.FormatMessage(exctx, level, msg)
	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) E - Msg test"

	match, _ := regexp.MatchString(expectedRegex, fmtedMsg)

	if !match {
		t.Errorf("Test failed, '%s' has incorrect format", fmtedMsg)
	}
}

func TestDefaultFormatterFatal(t *testing.T) {
	exctx := exctx.Create()
	level := Fatal
	msg := "Msg test"

	f := NewDefaultFormatter()

	fmtedMsg := f.FormatMessage(exctx, level, msg)
	expectedRegex := "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}.\\d{2,8}Z \\(\\w{8}\\) F - Msg test"

	match, _ := regexp.MatchString(expectedRegex, fmtedMsg)

	if !match {
		t.Errorf("Test failed, '%s' has incorrect format", fmtedMsg)
	}
}
