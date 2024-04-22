package main

import "fmt"

// 構造体
type Person struct {
	Name string
	Age  int
}

// Personのメソッド
func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func main() {
	person := Person{Name: "John", Age: 30}
	person.SayHello()
}
