package main

import "fmt"

// Maps are used to store data values in key:value pairs.
func main() {
	// Create Maps Using var and :=
	var a = map[string]string{}
	// "brand": "Ford", "model": "Mustang", "year": "1964"
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	// Create Maps Using the make() Function:
	c := make(map[string]int)
	c["Oslo"] = 1
	c["Bergen"] = 2
	c["Trondheim"] = 3
	c["Stavanger"] = 4
	fmt.Printf("c\t%v\n", c)

}
