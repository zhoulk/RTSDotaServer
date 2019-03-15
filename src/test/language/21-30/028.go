package main

// func main() {
// 	c1 := make(chan string, 1)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		c1 <- "one"
// 	}()

// 	select {
// 	case msg := <-c1:
// 		fmt.Println("receive ", msg)
// 	case <-time.After(time.Second * 1):
// 		fmt.Println("timeout ")
// 	}

// 	c2 := make(chan string, 1)

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		c2 <- "two"
// 	}()

// 	select {
// 	case msg := <-c2:
// 		fmt.Println("receive ", msg)
// 	case <-time.After(time.Second * 3):
// 		fmt.Println("timeout ")
// 	}
// }
