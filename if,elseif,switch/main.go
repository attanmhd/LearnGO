package main

import "fmt"

func main() {

	title:= "Golang the best language"

	for index,letter := range title {
		
		letterString := string(letter)

			if index %2 == 1 {
				switch letterString {
					case "a","i","u","e","o" :
						fmt.Println("index ke :" ,index,letterString)
					}
				}
	}

}