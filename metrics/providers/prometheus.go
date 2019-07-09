package metrics

import (
	"github.com/hugoluchessi/gobservable/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusCounter struct {
	cv  *prometheus.CounterVec
	lvs metrics.LabelValues
}

func NewPrometheusCounter(metricName string, labelNames ...string) *PrometheusCounter {
	cv := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: metricName,
	}, labelNames)

	return &PrometheusCounter{
		cv:  cv,
		lvs: labelNames,
	}
}

func (c *PrometheusCounter) With(labelValues ...string) metrics.Counter {
	return &PrometheusCounter{
		cv:  c.cv,
		lvs: c.lvs.With(labelValues...),
	}
}

func (c *PrometheusCounter) Add(delta float64) {
	c.cv.With(makeLabels(c.lvs...)).Add(delta)
}

func makeLabels(labelValues ...string) prometheus.Labels {
	labels := prometheus.Labels{}
	for i := 0; i < len(labelValues); i += 2 {
		labels[labelValues[i]] = labelValues[i+1]
	}
	return labels
}
