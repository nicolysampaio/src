package main

import "fmt"

func main() {
	name := "Nicoly"
	version := 1.1

	fmt.Println("Hello, Mr/Ms.", name)
	fmt.Println("This program is in version", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Display logs")
	fmt.Println("0 - Exit the program")

	var option int
	fmt.Scan(&option)

	fmt.Println("The chosen option was", option)

	switch option {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying logs...")
	case 0:
		fmt.Println("Exiting program...")
	default:
		fmt.Println("Invalid command")
	}
}
