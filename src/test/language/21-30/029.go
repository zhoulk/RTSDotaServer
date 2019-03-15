package main

// func main() {
// 	messages := make(chan string, 2)
// 	signals := make(chan bool)

// 	// go func() {
// 	// 	messages <- "hi"
// 	// }()

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("receive ", msg)
// 	default:
// 		fmt.Println("no message receive")
// 	}

// 	//messages <- "hi"

// 	msg := "hi"
// 	select {
// 	case messages <- msg:
// 		fmt.Println("sent message", msg)
// 	default:
// 		fmt.Println("no message sent")
// 	}

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("received message", msg)
// 	case sig := <-signals:
// 		fmt.Println("received signal", sig)
// 	default:
// 		fmt.Println("no activity")
// 	}
// }
