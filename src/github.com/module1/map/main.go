package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
		"funcB": func() int {
			return 2
		},
	}

	fmt.Println(myFuncMap)

	f := myFuncMap["funcA"]
	fmt.Println(f())

	value, exits := myMap["a"]
	if exits {
		fmt.Println(value)
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}
