package main


import "fmt"

// var arr [5]int
// var arr2 [5][5]int

// 声明时初始化
// var arr = [5]int{1,2,3,4,5}
// arr := [5]int{1,2,3,4,5}

func main(){
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
}

