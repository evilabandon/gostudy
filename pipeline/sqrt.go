package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	for _, n := range nums {
		out <- n
	}
	close(out)
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	c := sq(sq(gen(2, 3)))
	for out := range c {
		fmt.Println(out)
	}

	in := gen(2, 3, 4)
	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in)
	for n := range c1 {
		fmt.Println("c1", n)
	}
	for n := range c2 {
		fmt.Println("c2", n)
	}
	for n := range c3 {
		fmt.Println("c3", n)
	}

}
