package main

import (
	"fmt"
)

// Silahkan kerjakan latihan pada soal berikut :

// Buatlah sebuah struct dengan nama Student. Struct ini mempunyai properti Name dan Class. Lalu, dia memiliki 2 method yaitu SetMyName(name string) dan CallMyName().

// SetMyName(name string) harus bisa melakukan perubahan pada nama si student
// CallMyName() akan menampilkan Hello, My Name is <nama kamu>.

// JAWAB
type Student struct {
	Name  string
	Class string
}

func (c *Student) setMyName(name string, class string) {
	c.Name = name
	c.Class = class
}

func (c Student) CallMyName() {
	fmt.Println("Hello, My Name is ", c.Name+" and class", c.Class)
}

func main() {
	school := Student{Name: "Udin", Class: "IPS 5"}

	school.setMyName("Aryo", "IPA 2")

	school.CallMyName()
}
