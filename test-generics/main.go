package main

import (
	"fmt"
	"reflect"
)

func main() {
	t1 := 123
	fmt.Printf("output: %v (type: %s)\n", plusOne(t1), reflect.TypeOf(plusOne(t1)))

	t2 := 123.45
	fmt.Printf("output: %v (type: %s)\n", plusOne(t2), reflect.TypeOf(plusOne(t2)))

	fmt.Printf("output: %v (type: %s)\n", sum(t1, t1), reflect.TypeOf(sum(t1, t1)))
	fmt.Printf("output: %v (type: %s)\n", sum(t2, t2), reflect.TypeOf(sum(t2, t2)))
}

func plusOne[V int | int16 | int32 | int64 | float32 | float64](t V) V {
	return t + 1
}

func sum[V int | int16 | int32 | int64 | float32 | float64](t1, t2 V) V {
	return t1 + t2
}
