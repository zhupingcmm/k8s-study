package main

import (
	"context"
	"fmt"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	go func(c context.Context) {
		fmt.Println("====", c.Value("a"))
	}(ctx)

	//timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	//defer cancel()
}
