package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtsWorkers := 100

	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 100000; i++ {
		ch <- i
	}
}
