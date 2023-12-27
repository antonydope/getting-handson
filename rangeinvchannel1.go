package main

import (
	"fmt"
	"time"
)

func testSend(ch chan int) {
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
	close(ch)
}
func main(){
	ch := make(chan int)
	go testSend(ch) 
	for value := range ch {
		fmt.Println("Received:",value)
	}
	fmt.Println("channel closed")
}