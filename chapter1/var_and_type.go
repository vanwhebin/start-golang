package chapter1

import "fmt"


var a int  // 如果没有赋值默认为0
var b int = 1  // 声明时赋值
var c = 1 // 声明时赋值
var d int
var e string

// 语法糖
//d := 1
//e := "hello world"

// 简单类型
/*
- nil 空值
- int (int8 int16 int32 int64) 整型
- float (float32 float64)
- byte (等价与uint8)
- string 字符串类型
- boolean （true or false） 布尔值类型
*/

var a1 int8 = 10
var c1 byte = 'a'
var b1 float32 = 12.2
var msg = "hello world"
var ok bool

//ok := false

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
	fmt.Println(msg)
	fmt.Println(ok)
}



