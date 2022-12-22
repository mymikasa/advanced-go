package ch01

import (
	"fmt"
	"testing"
)

func TestGenerateNatural(t *testing.T) {
	ch := GenerateNatural()

	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)
	}
}

func TestChannel(t *testing.T) {
	ch := GenerateNatural()
	prime := <-ch
	fmt.Println(prime)

	ch = PrimeFilter(ch, prime)
	ch = PrimeFilter(ch, 3)

	for i := 0; i < 10; i++ {
		prime = <-ch
		fmt.Println(prime)
	}

}
