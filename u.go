package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// create a slice of int from 1 to 10
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := 0
	maxLen := len(i)

	for len(i) != 0 {
		fmt.Println("current Index", b)
		fmt.Println("current slice", maxLen)
		fmt.Println("current numbers")
		fmt.Println(i)

		if maxLen == 1 {
			i = []int{}
			continue
		}

		if b < maxLen {
			b = 0
		}

		randomNumber := rand.Intn(10)

		if randomNumber%2 == 0 {
			rIndex := b
			fmt.Println("remove index", b, "value", i[b])
			i = append(i[:rIndex], i[rIndex+1:]...)
			maxLen -= 1
		}

		b += 1

		fmt.Println()
	}
}
