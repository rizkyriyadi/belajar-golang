// Parameter tidak wajib, tapi bisa dimasukkan lebih dari satu

package main

import "fmt"

func dataDiri(nama string, umur int) {
	fmt.Println("Nama =", nama, "Umur =", umur)
}

func main() {
	// 		  Nama 			Umur
	dataDiri("Rizky Riyadi", 19)
}
