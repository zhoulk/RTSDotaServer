package main

// func main() {
// 	var state = make(map[int]int)

// 	var lock = &sync.Mutex{}

// 	var ops uint64 = 0

// 	for i := 1; i <= 100; i++ {
// 		total := 0
// 		go func() {
// 			key := rand.Intn(5)
// 			lock.Lock()
// 			total += state[key]
// 			lock.Unlock()
// 			atomic.AddUint64(&ops, 1)
// 			runtime.Gosched()
// 		}()
// 	}

// 	for w := 0; w < 10; w++ {
// 		go func() {
// 			for {
// 				key := rand.Intn(5)
// 				val := rand.Intn(100)
// 				lock.Lock()
// 				state[key] = val
// 				lock.Unlock()
// 				atomic.AddUint64(&ops, 1)
// 				runtime.Gosched()
// 			}
// 		}()
// 	}

// 	time.Sleep(time.Second)

// 	opsFinal := atomic.LoadUint64(&ops)
// 	fmt.Println("ops:", opsFinal)

// 	lock.Lock()
// 	fmt.Println("state:", state)
// 	lock.Unlock()
// }
