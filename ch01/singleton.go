package ch01

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance_v1() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}
