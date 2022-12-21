package ch01

import (
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

func worker01(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

var total2 uint64

func worker02(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 10; i++ {
		atomic.AddUint64(&total2, i)
	}
}
