package main

// func main() {
// 	messages := make(chan string)

// 	go func() {
// 		fmt.Println("gogoggooo")
// 		messages <- "ping"
// 		fmt.Println("gogoggooo222")
// 	}()

// 	go func() {
// 		fmt.Println("gogoggooo333")
// 		messages <- "ping2"
// 		fmt.Println("gogoggooo4444")
// 		messages <- "ping3"
// 	}()

// 	msg := <-messages
// 	fmt.Println(msg)
// 	msg = <-messages
// 	fmt.Println(msg)
// 	msg = <-messages
// 	fmt.Println(msg)
// }
