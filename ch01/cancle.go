package ch01

import (
	"context"
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, cancel chan bool) {

	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")

		case <-cancel:
			return
			// 退出
		}
	}
}

func ctx_worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
