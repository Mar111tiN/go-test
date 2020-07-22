package main

import (
    "fmt"
    "example.com/martin/test/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("Hello, world."))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
