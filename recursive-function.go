// Memanggil dirinya sendiri

package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	fmt.Println(factorial(5))
}
