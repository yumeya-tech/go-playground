package main

import (
	"fmt"
	"os"
)

func main() {
	// ファイルの作成
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 関数終了時にリソースを閉じる
	defer file.Close()

	// ファイルへの書き込み
	_, err = file.WriteString("Hello, file world!\n")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
