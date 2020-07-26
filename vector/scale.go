package vector

import (
	"fmt"
	"math"

	"github.com/Mar111tiN/go-test/add"
)

// Abs Vertex method returning absolute value of 2D-vector
func (v *Vector2D) Abs() float64 {
	return math.Sqrt(add.Sq(v.X) + add.Sq(v.Y))
}

// Abs Vector method returning absolute value of vector
func (v *Vector3D) Abs() float64 {
	return math.Sqrt(add.Sq(v.X) + add.Sq(v.Y) + add.Sq(v.Z))
}

// Abs Vertex method returning absolute value of 1D-vector
func (v *MyFloat) Abs() float64 {
	if *v < 0 {
		return float64(-*v)
	}
	return float64(*v)
}

// Scale expands V and returns the new absolute value
// Scale uses an interface type V that implements Abs() float
func Scale(v interface {
	Abs() float64
}, f float64) float64 {
	switch v := v.(type) {
	case *Vector3D:
		v.X = v.X * f
		v.Y = v.Y * f
		v.Z = v.Z * f
	case *Vector2D:
		v.X = v.X * f
		v.Y = v.Y * f
	// if v is of type *MyFloat, a conversion to float64 has to be performed prior
	// to multiplication and assignment of its pointer to v
	case *MyFloat:
		x := MyFloat(float64(*v) * f)
		v = &x
	}
	// test Vector using Type assertion
	if z, ok := v.(*Vector3D); ok {
		fmt.Println("Z-value of vector", z)
	} else {
		fmt.Println("Parameter Z is not available")
	}
	return v.Abs()
}
