package main

import (
	"fmt"
	"time"
)

func producer(ch chan string, name string , interval time.Duration) {
	for i := 1; i++ {
		time.Sleep(interval)
		ch <- fmt.Sprintf("%s: %d", name, i)
	}
}

func fanIn(ch1,ch2 <- chan string) <- chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case s := <- ch1:
				out <- s
			case s := <- ch2:
				out <- s
			}
		}
	}()
	return out
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1, "Producer 1", 100*time.Millisecond)
	go producer(ch2, "Producer 2", 250*time.Millisecond)

	merged := fanIn(ch1, ch2)

	for i := 0; i < 10; i++ {
		fmt.Println(<-merged)
	}
}

