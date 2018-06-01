package exctx

import (
	"time"
	"github.com/google/uuid"
	"testing"
	"github.com/hugoluchessi/gotoolkit/clock"
)

func TestNewExecutionContext(t *testing.T) {
	tm := time.Now()
	clock := clock.NewClock()
	clock.SetMockNow(tm)

	ctx := NewExecutionContext(clock)

	if ctx == nil {
		t.Error("[ctx] cannot be nil.")
	}
}
func TestNewExecutionContextNilClock(t *testing.T) {
	defer func() {
        if r := recover(); r == nil {
            t.Errorf("nil value on [clock] must panic.")
        }
	}()
	
	_ = NewExecutionContext(nil)
}

func TestNewWithTransaction(t *testing.T) {
	tm1 := time.Now()
	time.Sleep(1 * time.Microsecond)
	tm2 := time.Now()
	clock := clock.NewClock()
	clock.SetMockNow(tm1)
	uuid := uuid.New()

	ctx := NewWithTransaction(clock, uuid, tm2)

	if ctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	if ctx.ID.String() != uuid.String() {
		t.Errorf("[ctx] has invalid uuid, expected '%s' got '%s'.", uuid.String(), ctx.ID.String())
	}

	if ctx.TStarted.UnixNano() != tm2.UnixNano() {
		t.Errorf("[ctx] has invalid transaction started, expected '%d' got '%d'.", ctx.TStarted.UnixNano(), tm2.UnixNano())
	}

	if ctx.CStarted.UnixNano() != tm1.UnixNano() {
		t.Errorf("[ctx] has invalid transaction started, expected '%d' got '%d'.", ctx.CStarted.UnixNano(), tm1.UnixNano())
	}
}

func TestNewWithTransactionNilClock(t *testing.T) {
	defer func() {
        if r := recover(); r == nil {
            t.Errorf("nil value on [clock] must panic.")
        }
	}()

	tm := time.Now()
	uuid := uuid.New()

	_ = NewWithTransaction(nil, uuid, tm)
}