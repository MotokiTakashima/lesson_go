package main

import "fmt"

// 変数

// 関数の外で暗黙的な定義はできない
// i5 := 500
// 明示的な定義は可能
var i5 int = 500

// 関数の呼び出し
func outer() {
	var s4 string = "outer!!"
	fmt.Println(s4)
}

func main() {
	// 変数の定義　var 型　= 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello golang"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Go!!!"
	)
	fmt.Println(i2, s2)

	// 型の定義のみ
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 暗黙的な定義（型の指定がない）
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)

	// i4 := 450
	// fmt.Println(i4)

	// i4 = "Golang"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// 定義のみではエラーになる
	// var s5 string = "Not Use"
}
