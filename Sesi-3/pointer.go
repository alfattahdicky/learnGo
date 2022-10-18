package main

import "fmt"

func main()  {
	// pointer => variable spesial untuk menyimpan alamat memori variable lainnya
	// reference atau alamat yang sama itu membuat saling berhubungan 


	// pendahuluan 
	// var firstNumber *int
	// var secondNumber *int

	// memory address => menggunakan &
	// mengambil alamat memori dari variable menggunakan &
	// untuk mendapatkan nilai asli dari variable pointer menggunakan *

	// var firstNumber int = 4
	// var secondNumber *int = &firstNumber
	
	// fmt.Println("firstnumber value =", firstNumber)
	// fmt.Println("firstnumber memori address =", &firstNumber)

	// fmt.Println("secondnumber value =", *secondNumber)
	// fmt.Println("secondnumber memori address =", secondNumber)


	// changing value through pointer
	// memori tidak berubah tetapi nilai yang berubah
	var firstPerson string = "john"
	var secondPerson *string = &firstPerson

	fmt.Println(*secondPerson) // john
	fmt.Println(secondPerson)
	*secondPerson = "Doe"

	fmt.Println(*secondPerson) // doe
	fmt.Println(secondPerson)

	// pointer as parameter
	var a int = 10

	fmt.Println(a) // 10
	changeValue(&a)
	fmt.Println(a) // 20
}

func changeValue(number *int)  {
	*number = 20
}