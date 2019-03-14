package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[1])
	fmt.Println(len(a))

	b := [5]int{}
	fmt.Println(b)

	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println(d)
}
