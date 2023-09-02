package main

import "fmt"

func main() {
	myArray := [6]int{1, 2, 3, 4, 5, 8}
	mySlice := myArray[1:3]
	fmt.Println(mySlice)
}
