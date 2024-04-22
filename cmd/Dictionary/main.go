package main

import "fmt"

func main() {
	// キーがstring、値がstringのmapを生成
	var colors = map[string]string{"red": "#ff0000", "green": "#00ff00"}
	// mapに値を追加
	colors["blue"] = "#0000ff"
	// mapから値を取得
	fmt.Println(colors["blue"]) // #0000ff
}
