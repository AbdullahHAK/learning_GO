package main

import "fmt"

func main() {
	age := 25
	name := "john ibrar"
	height := 6.30922234

	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	//so println automatically adds space and also move Hello world to next line
	fmt.Println("Hello world")

	//whereas printf is same as it was in C and C++, boring
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type of age is %T\n", height)
	fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("name is %s\n", name)
}
