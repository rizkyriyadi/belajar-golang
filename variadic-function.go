package main

import "fmt"

// Hanya bisa di taro di akhir
func sumAll(numbers ...int) int {
	total := 0
	// "_" = indexnya dari 0 tapi ga di pake, value isi dari array numbers
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	ashiap := sumAll(10, 10, 21, 22, 33, 31, 31413)
	fmt.Println(ashiap)

	genap := []int{2, 4, 6, 8, 10}
	// Bisa di masukkan slice tinggal tambahin ...
	fmt.Println(sumAll(genap...))
}
