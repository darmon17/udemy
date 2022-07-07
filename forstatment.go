package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	slice := []string{"heri", "Sudarmono", "laki-laki"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	//slice atau array yang di pakai adalah array
	for _, value := range slice { //slice di gunakan range atau data yang akan di tampilkan
		fmt.Println("data", value)
	}

	//untuk map yang di pakai adalah key
	person := make(map[string]string)
	person["name"] = "eko"
	person["title"] = "newbie"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
