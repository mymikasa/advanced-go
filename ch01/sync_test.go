package ch01

import (
	"fmt"
	"sync"
	"testing"
)

func Test_worker(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(2)
	go worker01(&wg)
	go worker01(&wg)
	wg.Wait()

	fmt.Println(total.value)
}

var done = make(chan bool)

var msg string

func aGoroutine() {
	msg = "你好，世界"
	close(done)
}

func Test_channel(t *testing.T) {
	go aGoroutine()

	<-done
	println(msg)
}

func Test_sync(t *testing.T) {
	//var mu sync.Mutex
	//
	//mu.Lock()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("hello world")
	//	mu.Unlock()
	//}()
	//mu.Lock()
	//
	////mu.Unlock()

	//done := make(chan int, 10)
	//
	//for i := 0; i < cap(done); i++ {
	//	go func() {
	//		fmt.Println("hello world")
	//		done <- 1
	//	}()
	//}
	//
	//for i := 0; i < cap(done); i++ {
	//	<-done
	//}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {

		go func() {
			wg.Done()
			fmt.Println("hello world")
		}()
	}
	wg.Wait()
}
