package gttime

import (
	"time"
)

func Milliseconds(t time.Time) int64 {
	return milliseconds(t.UnixNano())
}

func ElapsedMilliseconds(s time.Time, e time.Time) int64 {
	ns := s.UnixNano()
	ne := e.UnixNano()

	return milliseconds(ne - ns)
}

func milliseconds(nano int64) int64 {
	return nano / int64(time.Millisecond)
}