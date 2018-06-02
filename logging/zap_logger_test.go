package logging

import (
	"time"
	"testing"
	"bufio"
	"bytes"
	"strings"

	"github.com/google/uuid"
	"github.com/hugoluchessi/gotoolkit/clock"
	"github.com/hugoluchessi/gotoolkit/exctx"	
)

func TestNewZapLogger(t *testing.T) {
	c := clock.NewClock()

	var b bytes.Buffer
    w := bufio.NewWriter(&b)

	cfg := LoggerConfig{ Debug, w }
	cfgs := []LoggerConfig{cfg}

	l := NewZapLogger(c, cfgs)

	if l == nil {
		t.Error("[l] must not be nil.")
	}
}

func createLoggingContext() (*exctx.ExecutionContext, *bytes.Buffer, *ZapLogger) {
	id := uuid.New()

	tm1 := time.Date(2018, time.January, 01, 12, 0, 0, 0, time.UTC)
	tm2 := tm1.Add(time.Millisecond * time.Duration(10))
	tm3 := tm2.Add(time.Millisecond * time.Duration(10))
	
	c := clock.NewClock()
	c.SetMockNow(tm2)
	ctx := exctx.NewWithTransaction(c, id, tm1)
	c.SetMockNow(tm3)

	var b bytes.Buffer
	cfg := LoggerConfig{ Debug, &b }
	cfgs := []LoggerConfig{cfg}

	l := NewZapLogger(c, cfgs)

	return ctx, &b, l
}

func TestDebug(t *testing.T) {
	ctx, b, l := createLoggingContext()

	l.Debug(
		ctx, 
		"Message test", 
		map[string]string { 
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := b.String()

	ell := "\"level\":\"debug\""
	ets := "\"ts\":"
	etid := "\"TID\":\"" + ctx.ID.String() + "\""
	etelapsed := "\"TElapsedMs\":\"20\""
	ecelasped := "\"CElapsedMs\":\"10\""
	etest1 := "\"test1\":\"value1\""
	etest2 := "\"test2\":\"value2\""

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, ets) {
		t.Error("Timestamp field not included.")
	}

	if !strings.Contains(content, etid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, etelapsed) {
		t.Error("Wrong transaction elapsed milliseconds.")
	}

	if !strings.Contains(content, ecelasped) {
		t.Error("Wrong current step elapsed milliseconds.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestInfo(t *testing.T) {
	ctx, b, l := createLoggingContext()

	l.Info(
		ctx, 
		"Message test", 
		map[string]string { 
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := b.String()

	ell := "\"level\":\"info\""
	ets := "\"ts\":"
	etid := "\"TID\":\"" + ctx.ID.String() + "\""
	etelapsed := "\"TElapsedMs\":\"20\""
	ecelasped := "\"CElapsedMs\":\"10\""
	etest1 := "\"test1\":\"value1\""
	etest2 := "\"test2\":\"value2\""

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, ets) {
		t.Error("Timestamp field not included.")
	}

	if !strings.Contains(content, etid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, etelapsed) {
		t.Error("Wrong transaction elapsed milliseconds.")
	}

	if !strings.Contains(content, ecelasped) {
		t.Error("Wrong current step elapsed milliseconds.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestWarn(t *testing.T) {
	ctx, b, l := createLoggingContext()

	l.Warn(
		ctx, 
		"Message test", 
		map[string]string { 
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := b.String()

	ell := "\"level\":\"warn\""
	ets := "\"ts\":"
	etid := "\"TID\":\"" + ctx.ID.String() + "\""
	etelapsed := "\"TElapsedMs\":\"20\""
	ecelasped := "\"CElapsedMs\":\"10\""
	etest1 := "\"test1\":\"value1\""
	etest2 := "\"test2\":\"value2\""

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, ets) {
		t.Error("Timestamp field not included.")
	}

	if !strings.Contains(content, etid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, etelapsed) {
		t.Error("Wrong transaction elapsed milliseconds.")
	}

	if !strings.Contains(content, ecelasped) {
		t.Error("Wrong current step elapsed milliseconds.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestError(t *testing.T) {
	ctx, b, l := createLoggingContext()

	l.Error(
		ctx, 
		"Message test", 
		map[string]string { 
			"test1": "value1",
			"test2": "value2",
		},
	)

	content := b.String()

	ell := "\"level\":\"error\""
	ets := "\"ts\":"
	etid := "\"TID\":\"" + ctx.ID.String() + "\""
	etelapsed := "\"TElapsedMs\":\"20\""
	ecelasped := "\"CElapsedMs\":\"10\""
	etest1 := "\"test1\":\"value1\""
	etest2 := "\"test2\":\"value2\""

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, ets) {
		t.Error("Timestamp field not included.")
	}

	if !strings.Contains(content, etid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, etelapsed) {
		t.Error("Wrong transaction elapsed milliseconds.")
	}

	if !strings.Contains(content, ecelasped) {
		t.Error("Wrong current step elapsed milliseconds.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

// func TestFatal(t *testing.T) {
// 	ctx, b, l := createLoggingContext()

// 	l.Fatal(
// 		ctx, 
// 		"Message test", 
// 		map[string]string { 
// 			"test1": "value1",
// 			"test2": "value2",
// 		},
// 	)

// 	content := b.String()

// 	ell := "\"level\":\"fatal\""
// 	ets := "\"ts\":"
// 	etid := "\"TID\":\"" + ctx.ID.String() + "\""
// 	etelapsed := "\"TElapsedMs\":\"20\""
// 	ecelasped := "\"CElapsedMs\":\"10\""
// 	etest1 := "\"test1\":\"value1\""
// 	etest2 := "\"test2\":\"value2\""

// 	if !strings.Contains(content, ell) {
// 		t.Error("Wrong log level.")
// 	}

// 	if !strings.Contains(content, ets) {
// 		t.Error("Timestamp field not included.")
// 	}

// 	if !strings.Contains(content, etid) {
// 		t.Error("Wrong transaction ID.")
// 	}

// 	if !strings.Contains(content, etelapsed) {
// 		t.Error("Wrong transaction elapsed milliseconds.")
// 	}

// 	if !strings.Contains(content, ecelasped) {
// 		t.Error("Wrong current step elapsed milliseconds.")
// 	}

// 	if !strings.Contains(content, etest1) {
// 		t.Error("Wrong test map params 1.")
// 	}

// 	if !strings.Contains(content, etest2) {
// 		t.Error("Wrong test map params 2.")
// 	}
// }