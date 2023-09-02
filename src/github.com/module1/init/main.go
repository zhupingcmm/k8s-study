package main

import (
	"fmt"
	_ "k8s/src/github.com/module1/init/a"
	_ "k8s/src/github.com/module1/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("  main function")
}
