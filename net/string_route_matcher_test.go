package net

import (
	"testing"
)

func TestNewStringRouteMatcher(t *testing.T) {
	rm := NewStringRouteMatcher("/something")

	if rm == nil {
		t.Error("Test failed, 'rm' is nil")
	}
}

func TestMatch(t *testing.T) {
	r := "/something"
	rm := NewStringRouteMatcher(r)
	match := rm.Match(r)

	if !match {
		t.Error("Test failed, route did not match")
	}
}

func TestNotMatch(t *testing.T) {
	r := "/something"
	r2 := "/something2"
	rm := NewStringRouteMatcher(r)
	match := rm.Match(r2)

	if match {
		t.Error("Test failed, different route match")
	}
}
