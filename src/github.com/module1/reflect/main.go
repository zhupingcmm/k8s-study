package main

import (
	"fmt"
	"reflect"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	myMap["c"] = "d"
	t := reflect.TypeOf(myMap)
	fmt.Println(t)
	v := reflect.ValueOf(myMap)
	fmt.Println(v)
	myStruct := T{A: "a"}
	v1 := reflect.ValueOf(myStruct)

	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Field %d: %v\n\n", i, v1.Field(i))
	}
	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}

	result := v1.Method(0).Call(nil)
	fmt.Println("result:", result)

}

type T struct {
	A string
}

func (t T) String() string {
	return t.A + "1"
}
