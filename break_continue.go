/*
Break : Menghentikan semua proses pengulangan

Continue : Skip ke iterasi selanjutnya
*/

package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i == 5 {
			// Kode setelah break ga bakal jalan lagi
			break
		}
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		// i yang di modulus 2 == 0 bakal di skip, jadinya bilangan ganjil doang >//<
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}

}
