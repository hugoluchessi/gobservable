package clock

import (
	"time"
	"testing"
)

func TestNewClock(t *testing.T) {
	clock := NewClock()

	if clock == nil {
		t.Error("[clock] cannot be nil.")
	}
}

func TestDefaultNow(t *testing.T) {
	clock := NewClock()

	if clock == nil {
		t.Error("[clock] cannot be nil.")
	}

	tm1 := clock.Now()
	time.Sleep(1 * time.Microsecond)
	tm2 := clock.Now()

	diff := tm2.UnixNano() - tm1.UnixNano()

	if diff <= 0 {
		t.Error("[tm2] must be greater than [tm1].")
	}
}

func TestSetMockNow(t *testing.T) {
	tm := time.Now()
	clock := NewClock()

	clock.SetMockNow(tm)
	
	if clock == nil {
		t.Error("[clock] cannot be nil.")
	}

	tm1 := clock.Now()
	time.Sleep(1 * time.Microsecond)
	tm2 := clock.Now()

	diff := tm2.UnixNano() - tm1.UnixNano()

	if diff != 0 {
		t.Error("[tm2] must be equal than [tm1].")
	}
}