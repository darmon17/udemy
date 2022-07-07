package main

import "fmt"

func getFullName() (string, string, int) {
	return "Heri", "Sudarmono", 23
}

func main() {
	firstName, lastName, age := getFullName()
	fmt.Println(firstName, lastName, age)
}
