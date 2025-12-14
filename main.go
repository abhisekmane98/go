package main

import (
	"fmt"

	"github.com/fyoder1821/go/add"
	"github.com/fyoder1821/go/substract"
)

func main() {
	fmt.Printf("Add: %d\n", add.Add(2, 3))
	fmt.Printf("Sub: %d\n", substract.Substract(2, 5))
}
