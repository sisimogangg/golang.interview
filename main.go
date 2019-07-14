package main

import (
	"fmt"
	"time"
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

func sumUp(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
		}()

		sum := 0
		for i := range in {
			sum += i
		}
		out <- sum

	}()
	return out
}

func main() {
	//done := make(chan struct{})

	out := gen()
	sum := sumUp(out)

	fmt.Println(<-sum)

	time.Sleep(4 * time.Second)
}
