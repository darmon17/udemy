package main

import "fmt"

func getFullName() (string, string, int) { //yang di minta returnnya tanpa parameter kembalian yang di minta 3 string, string, int
	return "Heri", "Sudarmono", 23 //return yang di kirim oleh func getFullName
}

func main() {
	firstName, lastName, age := getFullName() // data buat nampung yang di kirim oleh getFullName "Heri", "Sudarmono", 23
	fmt.Println(firstName, lastName, age)     //mencetak return yang di kirim oleh func getFullName
}
