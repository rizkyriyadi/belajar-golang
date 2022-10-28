/*
Satu satunya perulangan yang ada di Goland, tidak seperti bahasa lain
yang ada 2 atau lebih perulangan >//<
*/

package main

import "fmt"

func main() {

	i := 0

	// Di eksekusi terus sampai false

	for i <= 10 {
		if i%2 == 1 {
			fmt.Println(i)
		}
		// Di increment 1 supaya tidak terjadi infinite looping >//< bisa juga i ++
		i += 1
	}

	/*
		STATEMENT
		init ==> Statement sebelum for di eskekusi
		post ==> Selalu di eksekusi tiap akhir for
	*/
	// Genap := 1 ==> Sebagai init, dan genap++ ==> Sebagai post
	for genap := 1; genap <= 10; genap++ {
		if genap%2 == 0 {
			fmt.Println("Angka genap 1-10 = ", genap)
		}
	}

	// For pada array / Slice

	var hari = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	for j := 0; j < len(hari); j++ {
		fmt.Println("Index ke = ", j, "=", hari[j])
	}

	/*
		FOR RANGE

		for index, value := range (nama array / slice / map) {
			index menampung indexnya (index ke-0 sampai index ke-n)
			value menampung isi dari setiap indexnya
		}
	*/

	for index, value := range hari {
		fmt.Println("Index ke = ", index, "Berisi hari ", value)
	}
}
