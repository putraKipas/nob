package main

import (
	"fmt"
)

func main() {
	var car map[string]string
	car = map[string]string{}

	car["name"] = "BWM"
	car["color"] = "Black"

	// buat 2 buah fungsi :
	// 1 => fungsi yang mengembalikan sebuah string
	// pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna black

	// Jawaban
	fmt.Println(combineString(car["name"], car["color"]))

	// 2 => fungsi yang menampilkan hasil dari kembalian string
	// fungsi ini hanya bertugas untuk menampilkan kata

	// alur
	// simpan hasil dari return function kedalam sebuah variable message
	// tampilkan hasil dari variable message

	// Jawaban
	message := combineString(car["name"], car["color"])
	fmt.Println(message)

	// output => Mobil BMW berwarna Black
}

func combineString(str1 string, str2 string) string {
	return "Mobil " + str1 + " Berwarna " + str2
}
