package slice

type Scicer interface {
	Print(s string)
}
type IntSlice []int

type FloatSlice []float64

type SliceSlice []IntSlice

type Sl interface{}
