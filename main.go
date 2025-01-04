package main

import (
	"fmt"
	"learning_GO/greetings"
)

func main() {
	fmt.Println("trying to learn GO")
	greetings.Greet("its gonna be tough son")

	//basics of variables, can use both ways in first one we r getting specific but in second one it defines on runtime
	var name string = "orca"
	fmt.Println(name)
	var amount = 12
	fmt.Println(amount)
	//this one is used if we r getting value from function, another way
	person := "none"
	fmt.Println(person)

	//if u make name of variable or function with first letter capital
	//then u can directly export them and use them in other functions
	var Public = "this can be exported directly cuz first letter is capital"
	var private = "this can't be exported directly in other package"

	//rules apply for fucntions
	// func Public()  {
	// 	fmt.Println("this can be exported directly cuz first letter is capital")
	// }
	// func private(){
	// 	fmt.Println("this can't be exported directly in other package")
	// }

	fmt.Println(Public)
	fmt.Println(private)
	//thats why Println's P is capital. if u know u know
}
