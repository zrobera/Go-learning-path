package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(label string, reader *bufio.Reader) (string, error) {
	fmt.Print(label)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input, err
}

func calculateAverage(grades map[string]float64, numberOfSubjects int) float64 {
	var sum float64
	for _, value := range grades {
		sum += value
	}

	average := float64(sum)/float64(numberOfSubjects)
	return average
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, err := getInput("Enter name: ", reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	numberOfSubjectsStr, err := getInput("Enter number of subjects: ", reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Convert string to int
	numberOfSubjects, err := strconv.Atoi(numberOfSubjectsStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	for numberOfSubjects < 0 {
		fmt.Printf("Number of subject should be greater than 0 \n")
		numberOfSubjectsStr, err := getInput("Enter number of subjects: ", reader)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Convert string to int
		numberOfSubjects, err = strconv.Atoi(numberOfSubjectsStr)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
	}

	grades := make(map[string]float64)
	for i := 0; i < numberOfSubjects; i++ {
		prompt := fmt.Sprintf("Enter Subject %v: ", i+1)
		subject, err := getInput(prompt, reader)
		if err != nil {
			fmt.Print("Err:", err)
			return
		}
		prompt2 := "Enter grade: "
		gradeStr, err := getInput(prompt2, reader)
		if err != nil {
			fmt.Print("Err:", err)
			return
		}

		grade, err := strconv.ParseFloat(gradeStr,64)
		if err != nil {
			fmt.Print("Err:", err)
			return
		}
		for grade < 0 || grade > 100 {
			fmt.Printf("Grade should be in the range 0-100 \n")
			gradeStr, err := getInput(prompt2, reader)
			if err != nil {
				fmt.Print("Err:", err)
				return
			}
			grade, err = strconv.ParseFloat(gradeStr,64)
			if err != nil {
				fmt.Print("Err:", err)
				return
			}
		}
		grades[subject] = grade
		if err != nil {
			fmt.Print("Err:", err)
			return
		}
	}

	fmt.Println("Here is the result:")
	fmt.Printf("Your Name: %v \n", name)
	for key, value := range grades {
		fmt.Printf("%v:%v \n",key,value)
	}
	average := calculateAverage(grades, numberOfSubjects)
	fmt.Printf("Average: %v", average)
}
