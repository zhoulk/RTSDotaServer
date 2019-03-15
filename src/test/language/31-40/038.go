package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// type readOp struct {
// 	key  int
// 	resp chan int
// }

// type writeOp struct {
// 	key   int
// 	value int
// 	resp  chan bool
// }

// func main() {
// 	var ops uint64 = 0

// 	reads := make(chan *readOp)
// 	writes := make(chan *writeOp)

// 	go func() {
// 		var state = make(map[int]int)
// 		for {
// 			select {
// 			case read := <-reads:
// 				read.resp <- state[read.key]
// 			case write := <-writes:
// 				state[write.key] = write.value
// 				write.resp <- true
// 			}
// 		}
// 	}()

// 	for i := 1; i <= 100; i++ {
// 		go func() {
// 			for {
// 				read := &readOp{
// 					key:  rand.Intn(5),
// 					resp: make(chan int)}
// 				reads <- read
// 				<-read.resp
// 				atomic.AddUint64(&ops, 1)
// 			}
// 		}()
// 	}

// 	for i := 1; i <= 10; i++ {
// 		go func() {
// 			for {
// 				write := &writeOp{
// 					key:   rand.Intn(5),
// 					value: rand.Intn(100),
// 					resp:  make(chan bool)}
// 				writes <- write
// 				<-write.resp
// 				atomic.AddUint64(&ops, 1)
// 			}
// 		}()
// 	}

// 	time.Sleep(time.Second)

// 	opsFinal := atomic.LoadUint64(&ops)
// 	fmt.Println("ops:", opsFinal)
// }
