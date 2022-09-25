// Slice potongan dari array, Ukurannya bisa berubah (indexnya bisa bertambah)

/*
Membuat Slice dari Array
array[low:high] > Low index ke-0 dan high index len(array - 1) sebelum terkahir
array[:] > Dari index 0 sampai terkahir
*/

package main

import "fmt"

func main() {
	// [...] > No Limit, bisa pake [7]
	var hari = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	// Buat slice dari array hari
	// Mengambil 3 sampai Index terkahir

	var sliceSatu = hari[3:]
	fmt.Println(sliceSatu)

	/*
		Bila array atau slice dirubah, semua index akan ikut berubah.
		Karena terhubung satu sama lain
	*/

	// Membuat Slice baru dari index 5 sampai terkahir
	var sliceDua = hari[5:]
	fmt.Println(sliceDua)

	// Append, menambah ke index terkahir, kalo sudah penuh membuat array baru
	sliceDua = append(sliceDua, "MENAMBAH HARI BARU")
	fmt.Println(sliceDua)

	// Array awal tidak berubah karena Append di slice ke dua membuat array baru
	fmt.Println(hari)

	// Make Slice, nama = make([]tipedata, lenght, cap)
	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "Rizky"
	sliceBaru[1] = "Riyadi"

	fmt.Println(sliceBaru)
	fmt.Println("Panjang atau len sliceBaru =", len(sliceBaru))
	fmt.Println("Kapasitas atau cap sliceBaru =", cap(sliceBaru))

	sliceBaru = append(sliceBaru, "Mr Yadi")

	fmt.Println(sliceBaru)
	fmt.Println("Panjang atau len sliceBaru =", len(sliceBaru))
}
