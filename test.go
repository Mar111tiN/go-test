package main

import (
	"fmt"
	"math"

	"github.com/Mar111tiN/go-test/add"
)

// MyFloat helper type
type MyFloat float64

// Vector type for vectors
type Vector struct {
	X, Y, Z float64
}

//Vector2D type for 2D vector
type Vector2D struct {
	X, Y float64
}

// V interface implementing absolute values
type V interface {
	Abs() float64
}

// Abs Vector method returning absolute value of vector
func (v *Vector) Abs() float64 {
	return math.Sqrt(add.Sq(v.X) + add.Sq(v.Y) + add.Sq(v.Z))
}

// Abs Vertex method returning absolute value of 2D-vector
func (v *Vector2D) Abs() float64 {
	return math.Sqrt(add.Sq(v.X) + add.Sq(v.Y))
}

// Abs Vertex method returning absolute value of 1D-vector
func (v *MyFloat) Abs() float64 {
	if *v < 0 {
		return float64(-*v)
	}
	return float64(*v)
}

// Scale expands V and returns the absolute value
func Scale(v V, f float64) float64 {
	switch v := v.(type) {
	case *Vector:
		v.X = v.X * f
		v.Y = v.Y * f
		v.Z = v.Z * f
	case *Vector2D:
		v.X = v.X * f
		v.Y = v.Y * f
	case *MyFloat:
		x := MyFloat(float64(*v) * f)
		v = &x
	}
	return v.Abs()
}

// Scale expands vector and returns the absolute value
func (v *Vector2D) Scale(f float64) float64 {

	return v.Abs()
}

var v Vector

func main() {
	// create v as vector interface
	var v V
	// assign a 2D-Vector pointer
	v = &Vector2D{3, 4}
	fmt.Println(Scale(v, 3))

	// test Vector using Type assertion
	if z, ok := v.(*Vector); ok {
		fmt.Println(z)
	} else {
		fmt.Println("Not available")
	}
}
