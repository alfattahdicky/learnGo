package main

import "fmt"

func main()  {
	// alternative tipe data yang sudah ada pengecualian pada byte(uint8) dan rune(uint32)
	// menggunakan keyword type 
	type second = uint
	var hour second = 3600
	fmt.Printf("hour type: %T \n", hour)
}