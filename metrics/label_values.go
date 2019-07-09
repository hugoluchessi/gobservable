package metrics

type LabelValues []string

const (
	blankValue = "unknown"
)

func (lvs LabelValues) With(labelValues ...string) LabelValues {
	if len(labelValues)%2 != 0 {
		labelValues = append(labelValues, blankValue)
	}

	return append(lvs, labelValues...)
}
