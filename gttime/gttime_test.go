package gttime_test

import (
	"testing"
	"time"

	"github.com/hugoluchessi/gotoolkit/gttime"
)

func TestMilliseconds(t *testing.T) {
	t1 := time.Date(1999, 10, 4, 10, 30, 34, 0, time.Local)

	mili := t1.UnixNano() / int64(time.Millisecond)

	ms := gttime.Milliseconds(t1)

	if mili != ms {
		t.Errorf("Wrong milliseconds, expected %d got %d", mili, ms)
	}
}

func TestElapsedMilliseconds(t *testing.T) {
	t1 := time.Date(1999, 10, 4, 10, 30, 34, 0, time.Local)
	t2 := time.Date(1999, 10, 4, 10, 30, 34, int(300*time.Millisecond), time.Local)

	mili := int64(300)

	ms := gttime.ElapsedMilliseconds(t1, t2)

	if mili != ms {
		t.Errorf("Wrong miliseconds, expected %d got %d", mili, ms)
	}
}
