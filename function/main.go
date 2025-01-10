package main

import "fmt"

func simpleFunction() {
	fmt.Println("trying to learn GO, so that i can GO out of pakistan")
}

func add(a, b int) int {
	return a + b
}

//func add(a, b int) we wrote int with b only bcz func add(a int, b int, c int, d int, e int), so that we dont have to  do this
//func add(a, b int) int-> and this one indicates that function will return this type of value

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("start of my program")
	simpleFunction()
	var a int = 5
	var b int = 2
	ans := add(a, b)
	fmt.Println("sum is", ans)

	mul := multiply(a, b)
	fmt.Println("multiplication is", mul)
}
