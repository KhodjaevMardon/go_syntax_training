package main

import (
	"fmt"
)

func main() {
	n := 0
	oddsum, evensum := 0, 0
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			evensum += arr[i]
		} else {
			oddsum += arr[i]
		}
	}
	fmt.Printf("Sum of\nOdds : %d\nEvens : %d\n", oddsum, evensum)
}
