package main

import "fmt"

func main() {
	// Deklarasi Style satu
	var fullName [3]string

	fullName[0] = "Mohamad"
	fullName[1] = "Rizky"
	fullName[2] = "Riyadi"

	fmt.Println(fullName)

	fmt.Println(fullName[1])
	fmt.Println(fullName[2])

	// Style dua
	var angka = [3]int{
		21,
		21,
		13,
	}
	fmt.Println(angka)

	// Mendapatkan jumlah panjang array
	var PanjangAngka = len((angka))

	fmt.Println(PanjangAngka)

	// Langsung di print
	fmt.Println(len(fullName))

	// Mengubah Value
	fullName[2] = "Yadi"
	fmt.Println(fullName[2])
	fmt.Println(fullName)
}
