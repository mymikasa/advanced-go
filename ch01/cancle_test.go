package ch01

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}

func Test_ctx_worker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			err := ctx_worker(ctx, &wg)
			if err != nil {
				fmt.Println(err)
			}
		}()
	}

	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
