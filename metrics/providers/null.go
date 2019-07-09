package metrics

import "github.com/hugoluchessi/gobservable/metrics"

type NullCounter struct {
}

func (c *NullCounter) With(labelValues ...string) metrics.Counter {
	return c
}

func (c *NullCounter) Add(delta float64) {
	return
}
