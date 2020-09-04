package main

import (
	"fmt"
	"errors"
)

func main() {
	scores := []int {9,8,7,6,5,4,3}
	total := sum(scores)
	fmt.Println(total)

	result,err := calculation(90,4,"*")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

func calculation(number, numberTwo int, operation string) (int,error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unknown Operation")
	}

	return result,errorResult
}

func sum(numbers[]int) int {
	var total int
	for _,number := range numbers {
		total = total + number
	}
	return total
}