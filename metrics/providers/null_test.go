package metrics

import (
	"testing"
)

func TestWith(t *testing.T) {
	c := NullCounter{}

	c.With("test", "a")
}

func TestAdd(t *testing.T) {
	c := NullCounter{}

	c.Add(0.34)
}
