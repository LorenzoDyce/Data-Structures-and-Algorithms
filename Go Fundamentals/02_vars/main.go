package main

import "fmt"

func main() {
	// using var
	var name = "Lorenzo"
	var age int32 = 27
	var isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)
}
