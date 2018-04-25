package net

type StringRouteMatcher struct {
	r string
}

func NewStringRouteMatcher(route string) *StringRouteMatcher {
	return &StringRouteMatcher{route}
}

func (rm *StringRouteMatcher) Match(route string) bool {
	return rm.r == route
}
