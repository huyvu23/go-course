package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to my time study of go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// The Format method is used to format the time.Time value into a string representation.
	// The format string "01-02-2006" specifies the desired output format.
	// In this case, it formats the date as "MM-DD-YYYY".
	fmt.Println(presentTime.Format("01-02-2006"))

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2025, time.April, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))
}
