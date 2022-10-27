package main

import "fmt"

type mhs struct {
	npm   string
	nama  string
	kelas string
}
type krs struct {
	mhs
	kd_mk  string
	matkul string
	sks    int
}

func main() {
	var data_mhs = krs{}
	data_mhs.kd_mk = "IT009911"
	data_mhs.matkul = "Algoritma Pemrograman 3B"
	data_mhs.sks = 22

	fmt.Println("Masukkan NPM : ")
	fmt.Scan(&data_mhs.npm)
	fmt.Println("Masukkan Nama : ")
	fmt.Scan(&data_mhs.nama)
	fmt.Println("Masukkan Kelas : ")
	fmt.Scan(&data_mhs.kelas)

	fmt.Println("\n-----------------------")
	fmt.Println("NPM :", data_mhs.npm)
	fmt.Println("Nama :", data_mhs.nama)
	fmt.Println("Kelas :", data_mhs.kelas)
	fmt.Println("Kode Mata Kuliah :", data_mhs.matkul)
	fmt.Println("SKS :", data_mhs.sks)
	fmt.Println("\n-----------------------")

}
