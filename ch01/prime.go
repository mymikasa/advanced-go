package ch01

func GenerateNatural() chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			//fmt.Println(i)
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		for {
			if i := <-in; i%prime != 0 {
				//fmt.Printf("%v, %v\n", i, prime)
				out <- i
			}
		}
	}()

	return out
}
