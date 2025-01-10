package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero isn't possible")
	}
	return a / b, nil
}

func main() {
	fmt.Println("getting started with error handling")

	// ans := divide(10, 4)
	//so this line is saying divide func is returning two values but it is handling only one
	//so here we use "_" aka undrscore because we want to ignore or blacklist second value, like used below
	ans, _ := divide(10, 2)
	fmt.Println("division is: ", ans)

}
