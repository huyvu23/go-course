package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my slices study of go lang")
	// Slices are a flexible and powerful way to work with sequences of data in Go.
	// They are built on top of arrays and provide a more convenient and dynamic way to handle collections of elements.
	// Slices are more flexible than arrays because they can grow and shrink in size, while arrays have a fixed size.
	// Slices are a reference type, meaning that when you pass a slice to a function, you are passing a reference to the underlying array.
	// This allows for efficient sharing of data without copying the entire array.
	// Slices are created using the `make` function or by slicing an existing array or slice.
	// The `make` function is used to create a slice with a specified length and capacity.
	// The length of a slice is the number of elements it contains, while the capacity is the maximum number of elements it can hold.
	// When a slice is created, it has a length and capacity that are both set to the number of elements in the underlying array.
	// When you append elements to a slice, if the length exceeds the capacity, a new underlying array is created with double the capacity, and the elements are copied over.
	// This allows for efficient growth of the slice without the need to manually resize it.
	// Slices can be created using the `[]` operator, which creates a slice of the specified type.
	// The `len` function is used to get the length of a slice, and the `cap` function is used to get the capacity of a slice.
	// The `append` function is used to add elements to a slice, and it returns a new slice with the added elements.
	// The `copy` function is used to copy elements from one slice to another.
	// It takes two slices as arguments: the destination slice and the source slice.
	// The `copy` function copies elements from the source slice to the destination slice, up to the length of the shorter slice.
	// The `range` keyword is used to iterate over elements in a slice. It returns the index and value of each element in the slice.
	// The `for` loop is used to iterate over elements in a slice. It can be used with the `len` function to get the length of the slice.

	var fruitList = []string{"apple", "banana", "orange", "kiwi"}
	// fmt.Printf("Fruit List: %T", fruitList)

	// adding elements to a slice
	// fruitList = append(fruitList, "mango", "grape")
	// fmt.Println("\nFruit List after adding elements: ", fruitList)

	// `fruitList = append(fruitList[1:])` is slicing the `fruitList` slice starting from index 1 to the
	// end of the slice, and then appending the result back to the `fruitList` slice. This operation
	// effectively removes the first element from the `fruitList` slice.
	// fruitList = append(fruitList[1:])
	// fmt.Println("Fruit List after getting elements: ", fruitList)

	// `fruitList = append(fruitList[:2])` is slicing the `fruitList` slice from the beginning up to index 2
	fruitList = append(fruitList[:2])
	fmt.Println("Fruit List after append elements: ", fruitList)

	// Cắt ra một phần của slice fruitList — từ index 2 (bao gồm) đến index 3 (không bao gồm).
	// fruitList = append(fruitList[2:3])
	// fmt.Println("Fruit List after getting elements: ", fruitList)

	// ✅ 1. Dùng make với slice
	// Tạo slice []int có độ dài 3, capacity 5.
	//Các phần tử mặc định là 0.
	// Nếu chỉ truyền 2 đối số (make([]int, 3)), thì capacity = length.
	highScores := make([]int, 1)
	highScores[0] = 100
	// highScores[1] = 200 // panic: runtime error: index out of range [1] with length 1
	// Nhưng nếu dùng append thì không bị panic , nó sẽ tính và phân bổ lại bộ nhớ
	highScores = append(highScores, 200, 300, 400, 500)
	fmt.Println("length:", len(highScores), "capacity:", cap(highScores), highScores)

	// LOOP
	for i := 0; i < len(highScores); i++ {
		fmt.Println(highScores[i])
	}

	for i, itemHightScore := range highScores {
		fmt.Println("Index:", i, "itemHightScore:", itemHightScore)
	}

}
