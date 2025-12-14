package main

import (
	"fmt"

	"github.com/abhishekmane98/go/add"
	"github.com/abhishekmane98/go/substract"
)

func main() {
	fmt.Printf("Add: %d\n", add.Add(2, 3))
	fmt.Printf("Sub: %d\n", substract.Substract(2, 5))
}
