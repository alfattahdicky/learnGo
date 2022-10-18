package main

import "fmt"



func main()  {
	// empty interface => tipe data yang dapat menerima segala tipe data pada bahasa go
	var v interface{}

	v = 20
	// type assertion for reassign
	if value, ok := v.(int); ok == true {
		v = value * 9
	} 

	fmt.Printf("%#v \n", v)

	// map & slice with empty interface
	rs := []interface{}{1, "Dicky", true, 2}

	rm := map[string]interface{}{
		"name": "dicky",
		"age": 20,
		"status": true,
	}

	_,_ = rs, rm


}