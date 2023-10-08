// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU())
// 	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

// 	go speak(2, "Goroutine 1")
// }

// func main(){
// 	arr := map[string]string{
// 		"Name" : "NooBee",
// 		"Class" : "Backend Intermediate"
// 		"Address" : "Jakarta"
// 	}

// buatlah sebuah function print untuk nge handle hasil seperti dibawah
// pastikan menggunakan goroutine, agar urutan hasilnya itu bisa berbeda beda

// case 1
// Key : Name, Value : NooBee
// Key : Class, Value : Backend Intermediate
// Key : Addres, Value : Jakarta

// case 2
// Key : Class, Value : Backend Intermediate
// Key : Name, Value : NooBee
// Key