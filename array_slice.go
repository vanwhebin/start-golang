package main

import "fmt"

func main() {
	slice1 := make([]float32, 0)
	slice2 := make([]float32, 3, 5)

	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	slice2 = append(slice2, 1,2,3,4)
	fmt.Println(len(slice2), cap(slice2))

	// 子切片[start, end]

	sub1 := slice2[3:]
	sub2 := slice2[:3]
	sub3 := slice2[1:4]

	combined :=append(sub1, sub2...)
	fmt.Println(sub3)
	fmt.Println(combined)
}
