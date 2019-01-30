// Package keyrate wraps Limiter from golang.org/x/time/rate
// to perform key based limiting
package keyrate

import (
	"sync"

	"golang.org/x/time/rate"
)

// An IntLimiter controls how often events are allowed to happen per int key
type IntLimiter struct {
	m *sync.Mutex
	l map[int]*rate.Limiter

	n rate.Limit
	b int
}

// NewIntLimiter returns a new IntLimiter with up to rate r and burts of at most b
func NewIntLimiter(n rate.Limit, b int) *IntLimiter {
	return &IntLimiter{
		m: &sync.Mutex{},
		l: map[int]*rate.Limiter{},

		n: n,
		b: b,
	}
}

// Allow reports whether one event may happen at time now for the provided key
func (l *IntLimiter) Allow(key int) bool {
	l.m.Lock()

	if _, ok := l.l[key]; !ok {
		l.l[key] = rate.NewLimiter(l.n, l.b)
	}
	ok := l.l[key].Allow()

	l.m.Unlock()

	return ok
}

// A StringLimiter controls how often events are allowed to happen per string key
type StringLimiter struct {
	m *sync.Mutex
	l map[string]*rate.Limiter

	n rate.Limit
	b int
}

// NewStringLimiter returns a new StringLimiter with up to rate r and burts of at most b
func NewStringLimiter(n rate.Limit, b int) *StringLimiter {
	return &StringLimiter{
		m: &sync.Mutex{},
		l: map[string]*rate.Limiter{},

		n: n,
		b: b,
	}
}

// Allow reports whether one event may happen at time now for the provided key
func (l *StringLimiter) Allow(key string) bool {
	l.m.Lock()

	if _, ok := l.l[key]; !ok {
		l.l[key] = rate.NewLimiter(l.n, l.b)
	}
	ok := l.l[key].Allow()

	l.m.Unlock()

	return ok
}
