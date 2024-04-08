package main

import (
	"sync"
	"time"
)

type entry struct {
	createdAt time.Time
	value     []location
}

type cache struct {
	duration time.Duration
	entries  map[string]entry
	mu       sync.RWMutex
}

func newCache(dur time.Duration) *cache {
	return &cache{
		duration: dur,
		entries:  map[string]entry{},
	}
}

func (c *cache) add(key string, value []location) {
	c.mu.Lock()
	ent := entry{
		createdAt: time.Now(),
		value:     value,
	}
	c.entries[key] = ent
	c.mu.Unlock()

	time.AfterFunc(c.duration, func() { c.deleteEntry(key) })
}

func (c *cache) get(key string) []location {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.entries[key].value
}

func (c *cache) deleteEntry(key string) {
	c.mu.Lock()
	delete(c.entries, key)
	c.mu.Unlock()
}
