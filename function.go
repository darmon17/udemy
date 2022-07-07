package main

import "fmt"

func sayHallo() {
	fmt.Println("Say Hallo")
}

func main() {
	sayHallo()
	for i := 0; i < 5; i++ { // menentukan berapa kali di tampilkan
		sayHallo() //memanggil funcsi sayhallo
	}
}
