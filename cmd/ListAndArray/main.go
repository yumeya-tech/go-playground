package main

import "fmt"

func main() {
	// 配列
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[1]) // 2

	// スライス
	var slice []int = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println(slice[len(slice)-1]) // 6
}
