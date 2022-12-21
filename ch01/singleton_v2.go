package ch01

import "sync"

var (
	instance_v2 *singleton
	once        sync.Once
)

func Instance_V2() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
