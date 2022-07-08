package main

import "fmt"

func getFullname2() (firstName, lastName string, age int) { //declarasi nama variable yang akan di return
	firstName = "Heri"     //declarasi variable dan isi
	lastName = "Sudarmono" //declarasi variable dan isi
	age = 23               //declarasi variable dan isi

	return firstName, lastName, age // bisa di return semua variable
	//return bisa juga di kasih return gitu aja (sama dengan kirim semua variable yang di deklarasi di atas)

}

func main() {
	a, b, c := getFullname2()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
