package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey what's ur precious name?")
	// var name string

	// fmt.Scan(&name)
	// fmt.Println("nice to meet u Mr", name)
	// scan cant take after blank space

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
}
