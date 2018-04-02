package main

import (
	"fmt"
)

func main() {
	/**
	写入没读出，deadlock
	 */
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)


}
