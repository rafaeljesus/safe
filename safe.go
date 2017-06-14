package safe

import "sync"

// Safe wraps any value to a mutex.
type Safe struct {
	m     sync.RWMutex
	value interface{}
}

// New create a new Safe instance.
func New(value interface{}) *Safe {
	return &Safe{
		m:     sync.RWMutex{},
		value: value,
	}
}

// Get returns the value in a thread-safe fashion.
func (s *Safe) Get() interface{} {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.value
}

// Set sets a new value in a thread-save fashion.
func (s *Safe) Set(value interface{}) {
	s.m.Lock()
	defer s.m.Unlock()
	s.value = value
}
