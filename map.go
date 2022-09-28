// Map tidak seperti array, Map menggunakan Key-Value sebagai index dan bisa terus bertambah

package main

import "fmt"

func main() {

	dataDiri := map[string]string{
		"nama":        "Rizky Riyadi",
		"universitas": "Gunadarma",
	}

	dataDiri["alamat"] = "Jakarta"

	fmt.Println(dataDiri)
	fmt.Println(dataDiri["nama"])
	fmt.Println(dataDiri["alamat"])

	// Style 2, atau bisa dengan dataFisik := make(map[tipe_data]tipe_data)
	var dataFisik map[string]int = make(map[string]int)
	dataFisik["Umur"] = 19
	dataFisik["Tinggi"] = 175
	dataFisik["NIK"] = 2113
	fmt.Println(dataFisik)

	// Mengambil jumlah data di Map
	fmt.Println(len(dataFisik))

	// Menghapus data di Map dengan menyebutkan Key-Valuenya
	delete(dataFisik, "NIK")

	fmt.Println(dataFisik)
	fmt.Println(len(dataFisik))
}
