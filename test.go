package main

import (
	"fmt"

	"github.com/Mar111tiN/go-test/slice"
	"github.com/Mar111tiN/go-test/vector"
)

func main() {

	fmt.Println("=======VECTOR==============")
	v := vector.Vector2D{X: 3, Y: 4}
	fmt.Println(vector.Scale(&v, 3))

	// looking at slices
	fmt.Println("=======SLICES==============")
	a := slice.FloatSlice([]float64{0.3, 0.4, 0, 3.3, 1, 5})
	b := slice.IntSlice([]int{1, 2, 3, 4, 5})
	c1 := slice.IntSlice([]int{1, 2, 3})
	c2 := slice.IntSlice([]int{4, 5, 6})
	c3 := slice.IntSlice([]int{7, 8, 9})
	c := slice.SliceSlice([]slice.IntSlice{
		c1, c2, c3,
	})
	d := make(slice.IntSlice, 5)

	a.Print("a")
	b.Print("b")
	c.Print("c")
	d.Print("d")

	e := []byte("hello")
	e = append(e, " world"...)
	for i, s := range e {
		fmt.Printf("%v: %v\n", i, string(s))
	}
	fmt.Println("=======MAPS==============")
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	if value, ok := m["second"]; ok {
		fmt.Println("Value is ", value)
	}
	fmt.Println("=======FUNCS==============")
	// functional programming with go is possible but cumbersome
	compute := func(f func(int, int) int, x int, y int) int {
		return f(x, y)
	}
	fmt.Println("x+y", compute(func(x int, y int) int { return x + y }, 4, 5))
}
