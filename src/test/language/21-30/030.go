package main

// func main() {
// 	jobs := make(chan int, 5)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			j, more := <-jobs
// 			if more {
// 				fmt.Println("received job", j)
// 			} else {
// 				fmt.Println("no more chan")

// 				done <- true
// 				return
// 			}
// 		}
// 	}()

// 	for i := 1; i <= 13; i++ {
// 		jobs <- i
// 	}

// 	close(jobs)

// 	<-done
// }
