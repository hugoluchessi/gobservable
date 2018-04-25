package net

type RouteMatcher interface {
	Match(string) bool
}
