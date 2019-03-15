package main

import (
	"fmt"
	"sort"
)

type ByLenth []string

func (s ByLenth) Len() int {
	return len(s)
}

func (s ByLenth) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLenth) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	friuts := []string{"asasas", "asasa", "sas"}
	sort.Sort(ByLenth(friuts))
	fmt.Println(friuts)
}
