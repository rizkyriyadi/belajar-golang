package main

import "fmt"

// Harus di tulis tipe datanya walaupun sama
func myName() (string, string, string) {
	return "Mohamad", "Rizky", "Riyadi"
}

func main() {
	// Bisa di masukin ke variable
	firstName, middleName, lastName := myName()

	fmt.Println(firstName, middleName, lastName)

	// Pakai "_" bila tidak ingin di ambil
	nama, _, _ := myName()
	fmt.Println(nama)
}
