package main

import (
	"fmt"
)

func main()  {
	// string terbuat dari tipe data byte yaitu slice of bytes

	// looping over string (byte by byte)
	city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d \n", i , city[i]) // kode yang dikeluarkan ascii
	}

	// var cityByte []byte = []byte{74, 97, 107, 97, 114, 116, 97} 

	// fmt.Println(string(cityByte)) // jakarta


	// looping over string (rune by rune)
	// // rune => int32
	// str1 := "makan"
	// str2 := "mânca"

	// fmt.Printf("%d\n", len(str1))
	// fmt.Printf("%d\n", len(str2))

	// fmt.Printf("%d\n", utf8.RuneCountInString(str1))
	// fmt.Printf("%d\n", utf8.RuneCountInString(str2))

	// for i, s := range str1 {
	// 	fmt.Printf("index => %d, string => %s \n", i, string(s))
	// }

}