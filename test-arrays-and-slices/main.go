package main

import "fmt"

func main() {
	var arr1 []int = []int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Printf("len: %d, cap: %d\n", len(arr1), cap(arr1))

	arr2 := []int{1, 2, 3}
	arr2 = append(arr2, 4)
	fmt.Println(arr2)
	fmt.Printf("len: %d, cap: %d\n", len(arr2), cap(arr2))

	arr3 := []int{1, 2, 3}
	fmt.Printf("%p\n", arr3)
	arr3 = append(arr3, 4)
	fmt.Printf("%p\n", arr3)
	arr3 = append(arr3, 5)
	fmt.Printf("%p\n", arr3)
	arr3 = append(arr3, 6)
	fmt.Printf("%p\n", arr3)
	arr3 = append(arr3, 7)
	fmt.Printf("%p\n", arr3)
	fmt.Println(arr3)
	fmt.Printf("len: %d, cap: %d\n", len(arr3), cap(arr3))
}
