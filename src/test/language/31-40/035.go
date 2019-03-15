package main

// func main() {
// 	requests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		requests <- i
// 	}
// 	close(requests)

// 	limiter := time.NewTicker(time.Millisecond * 200)

// 	for j := range requests {
// 		<-limiter.C
// 		fmt.Println("handler ", j, time.Now())
// 	}

// 	burstyLimiter := make(chan time.Time, 3)
// 	for i := 1; i <= 3; i++ {
// 		burstyLimiter <- time.Now()
// 	}

// 	go func() {
// 		for t := range time.Tick(time.Millisecond * 200) {
// 			burstyLimiter <- t
// 		}
// 	}()

// 	burstyRequests := make(chan int, 5)
// 	for i := 1; i <= 5; i++ {
// 		burstyRequests <- i
// 	}
// 	close(burstyRequests)
// 	for req := range burstyRequests {
// 		<-burstyLimiter
// 		fmt.Println("request", req, time.Now())
// 	}
// }
