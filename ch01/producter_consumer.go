package ch01

import (
	"fmt"
	"time"
)

func Producter(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
