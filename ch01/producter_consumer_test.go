package ch01

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestProducter(t *testing.T) {
	ch := make(chan int, 64)

	go Producter(3, ch)
	go Producter(5, ch)

	go Consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("quit (%v)\n", <-sig)
}
