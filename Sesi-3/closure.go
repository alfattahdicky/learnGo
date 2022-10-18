package main

import (
	"fmt"
)


func main()  {
	// closure => anonymous function atau function tanpa nama disimpan ke dalam variable

	// ex closure
	// var evenNumbers = func(numbers ...int) []int  {
	// 	var result []int
	// 	for _, v := range numbers {
	// 		if v % 2 == 0 {
	// 			result = append(result, v)
	// 		}
	// 	}

	// 	return result
	// }

	// var numbers = []int{4, 93, 6, 20, 21, 23}
	// fmt.Println(evenNumbers(numbers...))

	// IIFE
	// var isPalindrome = func (str string) bool  {
	// 	var temp string = ""
	// 	for i := len(str) - 1; i >= 0; i-- {
	// 		temp += string(byte(str[i]))
	// 	} 
	// 	return temp == str
	// }("katak")

	// fmt.Println(isPalindrome)

	// closure as a return value

	// var studentList = []string {"dicky", "fattah", "azka"}

	// var find = findStudent(studentList)

	// fmt.Println(find("dicky"))

	// callback closure

	var find = findOddNumbers([]int {2,5,7,10}, func(number int) bool {
		return number % 2 != 0
	})

	fmt.Println(find)
}

// func findStudent(students []string) func(string)string  {
// 	return func(s string) string {
// 		var student string
// 		var position int

// 		for i, v := range students {
// 			if strings.ToLower(v) == strings.ToLower(s) {
// 				student = v
// 				position = i

// 				break
// 			}
// 		}

// 		if student == "" {
// 			return fmt.Sprintf("%s doesn't exist!", s)
// 		}

// 		return fmt.Sprintf("We found %s at position %d", s, position)
// 	}
// }

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers
}
