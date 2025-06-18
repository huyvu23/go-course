package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to my pointers study of go lang")
	x := 10
	p := &x         // p giữ địa chỉ của x
	fmt.Println(x)  // 10
	fmt.Println(*p) // 10 (lấy giá trị tại địa chỉ p)

	fmt.Println("address p: ", &p)

	// * operator is used to access the value at the memory address
	*p = 99        // thay đổi giá trị tại địa chỉ p
	fmt.Println(x) // 99 (x bị thay đổi)

	var y *int // y là một con trỏ trỏ đến kiểu int với giá trị là nil vì không được gán dữ liệu
	fmt.Println("address of y: ", &y)
	fmt.Printf("Type of x is : %T\n", y)
	fmt.Println("Value of y:", y)

	age := 23
	age2 := &age
	fmt.Println("address of age2:", age2)
	*age2 = 30
	fmt.Println("value of age2:", *age2)
	fmt.Println("value of age:", &age)
}
