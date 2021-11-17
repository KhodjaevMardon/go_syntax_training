package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	}
	runes := []rune(str)
	fmt.Println(runes)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, _ := reader.ReadString('\n')
	fmt.Println("Is your string palindrome: ", isPalindrome(strings.Trim(a, " \n")))
}
