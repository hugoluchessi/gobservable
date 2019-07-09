package metrics_test

import (
	"testing"

	"github.com/hugoluchessi/gobservable/metrics"
)

func TestWithTwoValues(t *testing.T) {
	lvs := metrics.LabelValues{}

	if len(lvs) != 0 {
		t.Fatal("FAIL: Initial length must be zero.")
	}

	value1 := "added only one value"
	value2 := "added only two values"

	nLvs := lvs.With(value1, value2)
	if nLvs[0] != value1 {
		t.Fatalf("FAIL: expected value '%s' got '%s'.", value1, nLvs[0])
	}

	if nLvs[1] != value2 {
		t.Fatalf("FAIL: expected value '%s' got '%s'.", value2, nLvs[1])
	}

	t.Log("SUCCESS")
}

func TestWithOneValue(t *testing.T) {
	lvs := metrics.LabelValues{}

	if len(lvs) != 0 {
		t.Fatal("FAIL: Initial length must be zero.")
	}

	value1 := "added only one value"
	blankValue := "unknown"

	nLvs := lvs.With(value1)
	if nLvs[0] != value1 {
		t.Fatalf("FAIL: expected value '%s' got '%s'.", value1, nLvs[0])
	}

	if nLvs[1] != blankValue {
		t.Fatalf("FAIL: expected value '%s' got '%s'.", blankValue, nLvs[1])
	}

	t.Log("SUCCESS")
}
