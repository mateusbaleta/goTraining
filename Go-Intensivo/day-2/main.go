package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data { // Reads channel values
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { // T1

	canal := make(chan int)

	qtWorkers := 50

	// SIMPLE LOAD BALANCER IMPLEMENTATION

	for i := 0; i < qtWorkers; i++ {
		go worker(i, canal)
	}

	//go worker(1, canal) // T2
	//go worker(2, canal) // T3

	for i := 0; i < 100; i++ {
		canal <- i
	}
}

// Below script is a brief explanation of process threading
// using Go routines and usage of channels for balancing
// memory usage for performance optimization
// ------------------------------------------------------------------------
//func contador(x int) {
//	for i := 0; i < x; i++ {
//		fmt.Println(i)
//		time.Sleep(time.Second)
//	}
//}
//
//func main() { // T1
//	// Basic example of channeling T1 routine output through T2
//	// Always note the importance of channel utilization for multi threading
//	// This prevents memory issues related to RACING
//	canal := make(chan string)
//
//	go func() { // T2
//		canal <- "opa"
//	}()
//
//	msg := <-canal
//	fmt.Println(msg) // Outputting T1
//
//	// Using go routines to enhance performance and multi thread operations
//	// !!  one of them can't receive the go command for effective  threading !!
//
//	go contador(10)
//	go contador(10)
//	contador(10)
//
//	// Process -> Alocates a memory block
//	// Thread -> Access memory block
//	// T1 - T2 -> Both access the same memory block
//	// Issue note -> Race Condition - The created  threads "fight" for the same allocated memory space
//	// AVOID RACE CONDITION AT ALL TIMES - This may impact the logical operation and results
//	// We share memory to communicate. We don't communicate sharing memory" - Rob Pike (a Go lang creator)
//
//	// Pipe results through CHANNELS between Go routines: #1 ROUTINE -> CHANNEL -> #2 ROUTINE
//	// #1 ROUTINE  GETS INPUT -> #2 ROUTINE GETS OUTPUT
//
//}
