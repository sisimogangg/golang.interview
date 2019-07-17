package main

import (
	"fmt"
)

func gen() <-chan int {
	out := make(chan int)

	go func() {
		defer func() {
			close(out)
		}()

		for i := 0; i < 10; i++ {
			out <- i
		}

	}()

	return out
}

func sumUp(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {

		sum := 0

		for i := range in {
			sum += i
		}

		out <- sum

	}()
	return out
}

func main() {
	sumCh := make(chan int)

	genCh := gen()
	// sum := sumUp(doneCh, genCh)

	defer func() {
		close(sumCh)
	}()

	go func() {

		sum := 0

		for i := range genCh {
			sum += i
		}

		sumCh <- sum

	}()
	//done <- struct{}{} // not required

	fmt.Println("Sum: ", <-sumCh)
	//time.Sleep(5 * time.Second) // only require if the sum func blocks

}
