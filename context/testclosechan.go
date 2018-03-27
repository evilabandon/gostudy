package main

import (
	"fmt"
)

/**
重复关闭 channel 会导致 panic。
向关闭的 channel 发送数据会 panic。
从关闭的 channel 读数据不会 panic，读出 channel 中已有的数据之后再读就是 channel 类似的默认值，比如 chan int 类型的 channel 关闭之后读取到的值为 0。
对于上面的第三点，我们需要区分一下：channel 中的值是默认值还是 channel 关闭了。可以使用 ok-idiom 方式，这种方式在 map 中比较常用。
 */
func main() {
	ch := make(chan int)
	go func(){
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for i:=range ch {
		fmt.Println(i)
	}



}
