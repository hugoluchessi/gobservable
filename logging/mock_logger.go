package logging

import (
	"bytes"
	"fmt"
	"io"
)

type MockLogger struct {
	b *bytes.Buffer
}

func NewMockLogger() *MockLogger {
	b := bytes.NewBufferString("")
	return &MockLogger{b}
}

func (l *MockLogger) Log(msg string, params map[string]interface{}) {
	mockLogMsg(l.b, Log, msg, flattenParams(params))
}

func (l *MockLogger) Error(msg string, params map[string]interface{}) {
	mockLogMsg(l.b, Error, msg, flattenParams(params))
}

func (l *MockLogger) String() string {
	return l.b.String()
}

func mockLogMsg(w io.Writer, lvl LogLevel, msg string, params []interface{}) {
	fmt.Fprintf(w, "level: %d ", lvl)
	fmt.Fprintf(w, "msg: %s ", msg)

	for i := 0; i < len(params); i += 2 {
		fmt.Fprintf(w, "%s: %s ", params[i], params[i+1])
	}
}
