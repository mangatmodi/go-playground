package main

import (
	"fmt"
	"math"
)

type mappable interface {
	able()
}

// Use reflection to find and init the
type from struct {
	i1 float32
}

func (*from) able() {}

type to struct {
	i int32
}

func (*to) able() {}

var input mappable = &from{3.14}
func main() {
	fmt.Println(m(round))
}

func round(i from) to {
	return to{int32(math.Floor(float64(i.i1)))}
}

func m(fn func(from) to) to {
	switch input.(type) {
	case *from:
		return fn(*input.(*from))
	default:
		panic("Wrong type")
	}
}
