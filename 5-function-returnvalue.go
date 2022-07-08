package main

import "fmt"

func getHallo(name string) string { // retrun yang diminta adalah string dan ada name sebagai parameter
	if name == "" { //kondisi nama kosong
		return "Hallo Bro" // eksekusi jika nama kosong
	} else {
		return "Hallo " + name //kondisi jika salah maka cetak Hallo + nama tidak pakai ,(koma)
	}
}

func main() {
	result := getHallo("Heri Sudamono") // result adalah variable getHallo adalah func, heri sudarmono sebagi return yang di minta/yang di kirim
	fmt.Println(result)                 // cetak data
	resultkosong := getHallo("")        //jika kosong maka yang di cetak adalah hallo bro
	fmt.Println(resultkosong)           // cetak variable resultkosong
}
