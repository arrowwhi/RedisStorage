package main

import (
	"sync"
	"time"
)

type Data struct {
	value     interface{}
	expiresAt time.Time
}

type Storage struct {
	store map[string]Data
	mu    sync.RWMutex
}

func NewStore() *Storage {
	return &Storage{
		store: make(map[string]Data),
	}
}

func (s *Storage) Set(key string, value interface{}, ttl time.Duration) {
	if ttl < 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	expiresAt := time.Now().Add(ttl)
	s.store[key] = Data{
		value:     value,
		expiresAt: expiresAt,
	}
	if ttl > 0 {
		time.AfterFunc(ttl, func() {
			s.Delete(key)
		})
	}
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.store[key]
	if !ok {
		return nil, false
	}
	return item.value, true
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}
