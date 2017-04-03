package main

import (
	"fmt"
	"golang.org/x/net/context"
)

func main() {
	bg := context.Background()
	fmt.Printf("bg = %v\n", bg)
}
