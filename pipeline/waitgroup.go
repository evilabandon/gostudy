package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println("start goroutine ", i)
			time.Sleep(2 * time.Second)
			fmt.Println("goroutine end ", i)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("all done")

}
