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

		select {
		case out <- sum:
		case <-done:
			break
		}

	}()
	return out
}

func main() {
	done := make(chan struct{})

	out := gen()
	sum := sumUp(done, out)
	//done <- struct{}{}

	fmt.Println(<-sum)
	//time.Sleep(5 * time.Second)

}
