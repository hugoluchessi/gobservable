package logging

type LogLevel int

const (
	Log LogLevel = iota
	Error
)

type Logger interface {
	Log(string, map[string]interface{})
	Error(string, map[string]interface{})
}
