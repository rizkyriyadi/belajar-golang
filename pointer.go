package main

import "fmt"

type yadi struct {
	nama, npm string
}

func main() {

	data1 := yadi{"Mohamad Rizky Riyadi", "50421837"}
	data2 := &data1

	*data2 = yadi{"Mamang", "2113"}

	fmt.Println(data1)

	// data2.nama = "yadi"

	// fmt.Println(data1)
	// fmt.Println(data2)

}
