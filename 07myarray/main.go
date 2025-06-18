package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to my array study of go lang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "orange"
	fruitList[3] = "mango"

	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Length of fruit list: ", len(fruitList))

	// list language programming
	programmingLanguages := [...]string{"go", "python", "java", "javascript"} // here length is inferred
	fmt.Println("Programming Languages: ", programmingLanguages)

	for index, language := range programmingLanguages {
		if language == "java" {
			programmingLanguages[index] = "c"
		}
		fmt.Println("language:", language)
	}
	fmt.Println("Programming Languages: ", programmingLanguages)
}
