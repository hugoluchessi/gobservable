package logging

import (
	"strings"
	"testing"
)

func TestNewMockLogger(t *testing.T) {
	l := NewMockLogger()

	if l == nil {
		t.Error("[l] must not be nil.")
	}
}

func TestMockLog(t *testing.T) {
	l := NewMockLogger()

	l.Log(
		"Message test",
		map[string]interface{}{
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := l.String()

	ell := "level: 0"
	etest1 := "test1: value1"
	etest2 := "test2: value2"

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestMockError(t *testing.T) {
	l := NewMockLogger()

	l.Error(
		"Message test",
		map[string]interface{}{
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := l.String()

	ell := "level: 1"
	etest1 := "test1: value1"
	etest2 := "test2: value2"

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}
