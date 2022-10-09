// Function yang mengembalikan Data

package main

import "fmt"

func getName(name string) string {
	return "Your name = " + name
}

func main() {
	// Karena mengembalikan bisa disimpan di variable
	nama := getName("Rizky")

	fmt.Println(nama)
	fmt.Println(getName("Yadi"))
}
