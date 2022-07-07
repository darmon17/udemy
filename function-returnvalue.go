package main

import "fmt"

func sayHallo(name string) string { // retrun yang diminta adalah string dan ada name sebagai parameter
	if name == "" { //kondisi nama kosong
		return "Hallo Bro" // ekse kusi jika nama kosong
	} else {
		return "Hallo " + name //kondisi jika salah maka cetak Hallo + nama tidak pakai ,(koma)
	}
}

func main() {
	result := sayHallo("Heri Sudamono") // result adalah variable sayHallo adalah func, heri sudarmono sebagi return yang di minta/yang di kirim
	fmt.Println(result)                 // cetak data
	resultkosong := sayHallo("")        //jika kosong maka yang di cetak adalah hallo bro
	fmt.Println(resultkosong)           // cetak variable resultkosong
}
