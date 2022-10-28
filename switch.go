package main

import "fmt"

func main() {

	kendaraan := "Terbang"

	switch kendaraan {
	case "Terbang":
		fmt.Println("Pesawat")

	case "Darat":
		fmt.Println("Mobil")

	case "Laut":
		fmt.Println("Perahu / Kapal Laut")

	// default sama seperti else pada Percabangan
	default:
		fmt.Println("Belum ada kendaraan seperti itu >//< ")
	}

}
