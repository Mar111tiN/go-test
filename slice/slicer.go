package slice

import (
	"fmt"
)

func (x *IntSlice) Print(s string) {
	fmt.Printf("IntSlice %s: len=%d cap=%d %v\n",
		s, len(*x), cap(*x), *x)
}

func (y *FloatSlice) Print(s string) {
	fmt.Printf("FloatSlice %s: len=%d cap=%d %v\n",
		s, len(*y), cap(*y), *y)
}

func (z *SliceSlice) Print(s string) {
	fmt.Printf("%s: Slice of IntSlices\n", s)
	for i := 0; i < len(*z); i++ {
		slice := *z
		fmt.Printf("%v -> len=%d cap=%d %v\n",
			i, len(slice[i]), cap(slice[i]), slice[i])
	}
}
