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

	var sportsList = []string{"soccer", "basketball", "tennis", "cricket"}
	fmt.Println("Sports List: ", sportsList)
	fmt.Println("Length of sports list: ", len(sportsList))

}
