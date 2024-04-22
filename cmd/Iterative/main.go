package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	numbers := []int{2, 3, 5, 7, 11}
	for i, num := range numbers {
		fmt.Printf("Index %d: %d\n", i, num)
	}
}
