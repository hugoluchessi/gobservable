package maps

import "testing"
func TestMergeStringMaps(t *testing.T) {
	key1 := "k1"
	key2 := "k2"
	value1 := "v1"
	value2 := "v1"

	a := map[string]string {
		key1: value1,
	}

	b := map[string]string {
		key2: value2,
	}

	c := MergeStringMaps(a, b)

	if c[key1] != a[key1] {
		t.Errorf("[c[key1]] must have the same value in map 'a', expected '%s' go '%s'.", a[key1], c[key1])
	}

	if c[key2] != b[key2] {
		t.Errorf("[c[key2]] must have the same value in map 'b', expected '%s' go '%s'.", b[key2], c[key2])
	}
}