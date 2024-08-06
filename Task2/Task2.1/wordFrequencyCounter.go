package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func wordFrequencyCount(words string) map[string]int{
	words = strings.ToLower(words)
	listWords := strings.Split(words, " ")
	var counter = make(map[string]int)


	for _, word := range listWords{
		var str string
		for _, ch := range word{
			if unicode.IsLetter(ch){
				str += string(ch)
			}
		}
		counter[str] += 1
	}
	return counter
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")

	input, _ := reader.ReadString('\n')
	frequencies := wordFrequencyCount(strings.TrimSpace(input))
	fmt.Println("Word Frequencies:", frequencies)
}