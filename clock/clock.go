package clock

import "time"

type funcTimeNow func() time.Time

type Clock struct{
	now funcTimeNow
}

func NewClock() *Clock {
	return &Clock{ 
		func () time.Time {
			return time.Now()
		},
	}
}

func (c *Clock) Now() time.Time {
	return c.now()
}

func (c *Clock) SetMockNow(t time.Time) {
	c.now = func () time.Time {
		return t
	}
}