package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 { // program akan otomatis berhenti jika kondisi = 5
			break
		}
		fmt.Println("break perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 { // untuk menampilkan data yang ganjil
			continue
		}
		fmt.Println("continue perulangan ke ", i)
	}
}
