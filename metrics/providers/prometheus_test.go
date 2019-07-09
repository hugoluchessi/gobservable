package metrics

import (
	"testing"
)

func TestNewPrometheusCounter(t *testing.T) {
	labelName := "label 1"
	labelValue := "label value 1"
	c := NewPrometheusCounter("test", labelName, labelValue)

	if c == nil {
		t.Fatalf("FAIL: should not return nil.")
	}

	if c.cv == nil {
		t.Fatal("FAIL: counter cannot be nil.")
	}

	if c.lvs[0] != labelName {
		t.Fatalf("FAIL: expected label name '%s' got '%s'.", c.lvs[0], labelName)
	}

	if c.lvs[1] != labelValue {
		t.Fatalf("FAIL: expected label value '%s' got '%s'.", c.lvs[1], labelValue)
	}

	t.Log("SUCCESS")
}

func TestNewPrometheusCounterWith(t *testing.T) {
	labelName := "label 1"
	labelValue := "label value 1"
	otherLabelName := "label 2"
	otherLabelValue := "label value 2"
	c := NewPrometheusCounter("test", labelName, labelValue)

	nc := c.With(otherLabelName, otherLabelValue)

	ntc, ok := nc.(*PrometheusCounter)

	if !ok {
		t.Fatalf("FAIL: wrongly typed counter.")
	}

	if ntc.lvs[0] != labelName {
		t.Fatalf("FAIL: expected label name '%s' got '%s'.", ntc.lvs[0], labelName)
	}

	if ntc.lvs[1] != labelValue {
		t.Fatalf("FAIL: expected label value '%s' got '%s'.", ntc.lvs[1], labelValue)
	}

	if ntc.lvs[2] != otherLabelName {
		t.Fatalf("FAIL: expected label name '%s' got '%s'.", ntc.lvs[0], labelName)
	}

	if ntc.lvs[3] != otherLabelValue {
		t.Fatalf("FAIL: expected label value '%s' got '%s'.", ntc.lvs[1], labelValue)
	}

	t.Log("SUCCESS")
}
