package main

// func main() {
// 	fmt.Println(time.Now())

// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(time.Second)
// 		c1 <- "one"
// 	}()

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		c2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println("received", msg1)
// 		case msg2 := <-c2:
// 			fmt.Println("received", msg2)
// 		}
// 	}

// 	fmt.Println(time.Now())
// }
