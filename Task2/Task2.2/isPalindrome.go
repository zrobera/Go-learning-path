package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPalindrome(input string) bool {
	
	var newString []rune
	for _, r := range input {
		if unicode.IsLetter(r) {
			newString = append(newString, unicode.ToLower(r))
		}
	}

	length := len(newString)
	for i := 0; i < length/2; i++ {
		if newString[i] != newString[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")

	input, _ := reader.ReadString('\n')
	isPalin := isPalindrome(strings.TrimSpace(input))
	if isPalin {
		fmt.Println("The input string is a palindrome.")
	} else {
		fmt.Println("The input string is not a palindrome.")
	}
}
