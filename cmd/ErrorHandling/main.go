package main

import (
	"errors"
	"fmt"
)

// 戻り値として error 型を返す。
func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		// エラー型を生成して返す
		return 0.0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(10.0, 0.0)
	if err != nil {
		// エラーがある場合
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
