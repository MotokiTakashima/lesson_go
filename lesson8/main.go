package main

import (
	"fmt"
)

//配列型

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素数を自動で決めてくれる
	arr4 := [...]string{"A", "S"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//配列の取り出し
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	//要素数が異なる場合、代入できない
	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)

	//配列の要素数を出力する
	fmt.Println(len(arr1))
}
