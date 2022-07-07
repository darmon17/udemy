package main

import "fmt"

func GetFullName() (string, string, int) {
	return "Heri", "Sudarmono", 23
}

func main() {
	firstName, lastName, age := GetFullName()
	fmt.Println(firstName, lastName, age)
}
