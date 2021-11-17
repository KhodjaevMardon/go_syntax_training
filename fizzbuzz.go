package main

import "fmt"

func main() {
	for i := 0; i < 66; i++ {
		if i%5 != 0 && i%3 != 0 {
			fmt.Print(i)
		} else {
			if i%5 == 0 {
				fmt.Print("Fizz")
			}
			if i%3 == 0 {
				fmt.Print("Buzz")
			}
		}
		fmt.Print("\n")
	}
}
