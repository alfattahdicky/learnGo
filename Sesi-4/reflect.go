package main

import (
	"fmt"
	"reflect"
)

func main()  {
	// reflect => melakukan inspeksi variabel yait mengambil informasi dari variable tersebut atau bahkan memanipulasinya
	// penting yg diketahui reflect.ValueOf() dan reflect.TypeOf()
	// reflect.valueof => mengembalikan objek berisi nilai pada variable
	// reflect.typeof => mengembalikan objek berisi tipe data variable

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variable", reflectValue.Int())
	// }

	fmt.Println("nilai variable", reflectValue.Interface())



}