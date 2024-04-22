package main

import "fmt"

func main() {
	// 変数
	var num int = 10
	// 定数
	const constNum int = 20

	fmt.Println(num)
	fmt.Println(constNum)

	// 型推論
	var typeInferenceBool = true  // bool型
	var typeInferenceNum = 10     // int型
	var typeInferenceStr = "abc"  // string型
	var typeInferenceFloat = 10.5 // float型

	fmt.Println(typeInferenceBool)
	fmt.Println(typeInferenceNum)
	fmt.Println(typeInferenceStr)
	fmt.Println(typeInferenceFloat)

	// 複数の変数を同時に宣言
	var x, y = 10, "Hello" // xにはint、yにはstringを代入

	fmt.Println(x, y)

	// 短い変数宣言
	shortNum := 20                // int型
	shortX, shortY := 30, "Hello" // int型とstring型

	fmt.Println(shortNum)
	fmt.Println(shortX, shortY)
}
