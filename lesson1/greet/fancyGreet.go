package greet

import "fmt"

func FancyGreet() {

	var choice string
	var booked [10]string

	fmt.Println("Welcome!")

	for choice != "0" {
		fmt.Println(
			"Choose what to do:\n" +
				"1 - book table\n" +
				"2 - show books\n" +
				"0 - exit",
		)
		fmt.Scanln(&choice)
		if choice != "1" && choice != "2" {
			fmt.Println("Wrong input. Please try again")
			continue
		}
		if choice == "1" {
			var name string
			var table int

			fmt.Println("Available tables:")
			for i := 0; i < len(booked); i++ {
				if booked[i] == "" {
					fmt.Printf("%d ", i+1)
				}
			}

			fmt.Print("\nEnter table number: ")
			fmt.Scanln(&table)

			for table < 1 || table > 10 || booked[table-1] != "" {
				fmt.Print("Table not found or already booked. Enter other table number: ")
				fmt.Scanln(&table)
			}

			fmt.Print("Enter your name: ")
			fmt.Scanln(&name)
			booked[table-1] = name
		} else if choice == "2" {
			fmt.Println("Books:")
			for i := 0; i < len(booked); i++ {
				if booked[i] == "" {
					fmt.Printf("%d - not booked\n", i+1)
				} else {
					fmt.Printf("%d - booked by %s\n", i+1, booked[i])
				}
			}
		}
	}
}
