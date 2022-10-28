package main

import "fmt"

func bio() (name, nickname string, age int) {
	name = "Mohamad Rizky Riyadi"
	nickname = "Rizky"
	age = 19

	return
}

func main() {
	nama, panggilan, umur := bio()

	fmt.Println(nama)
	fmt.Println(panggilan)
	fmt.Println(umur)
}
