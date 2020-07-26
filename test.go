package main

import (
	"fmt"

	"github.com/Mar111tiN/go-test/vector"
)

func main() {

	// assign a 2D-Vector pointer
	a := []float64{0.3, 0.4, 0, 3.3, 1, 5}
	v := vector.Vector2D{X: 3, Y: 4}
	fmt.Println(vector.Scale(&v, 3))

	fmt.Printf("Array is %v\n", a[1:5])
}
