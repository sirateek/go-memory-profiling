package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	key := "KEY"
	value := "HelloWorld"

	ctx = context.WithValue(ctx, key, value)

	result := ctx.Value(key)
	fmt.Println(result)

	ctx = context.WithValue(ctx, key, "Test2")
	result2 := ctx.Value(key)
	fmt.Println(result2)
}
