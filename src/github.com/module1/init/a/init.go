package a

import (
	"fmt"
	_ "k8s/src/github.com/module1/init/b"
)

func init() {
	fmt.Println("init from a")
}
