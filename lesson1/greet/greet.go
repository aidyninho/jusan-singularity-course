package greet

import "fmt"

func Greet() {

	var table int
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your table: ")
	fmt.Scanln(&table)
	fmt.Printf("Hello, %s! Your table is: %d\n", name, table)
}
