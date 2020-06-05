package main

import (
	"fmt"
	_ "fmt"
	"reflect"
	_ "reflect"
)

func main() {
	str1 := "Golang"
	str2 := "go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	fmt.Printf("%d %c\n", str2[2], str2[2])
	fmt.Println("len(str2): ", len(str2))
}
