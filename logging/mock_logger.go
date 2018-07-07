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
		key := params[i]
		value := params[i+1]

		switch value.(type) {
		case int:
			fmt.Fprintf(w, "%s: %d ", key, value)
		case int64:
			fmt.Fprintf(w, "%s: %d ", key, value)
		default:
			fmt.Fprintf(w, "%s: %s ", key, value)
		}
	}

	fmt.Fprintf(w, "\n")
}
