package main

import (
	"encoding/json"
	"fmt"
)

type MyType struct {
	Name string `json:"name"`
	Address
}
type Address struct {
	City string `json:"city"`
}

func main() {
	mt := MyType{Name: "test", Address: Address{City: "shanghai"}}
	b, _ := json.Marshal(&mt)
	fmt.Println(b)
}
