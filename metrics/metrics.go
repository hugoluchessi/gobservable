package metrics

type Counter interface {
	With(labelValues ...string) Counter
	Add(delta float64)
}
