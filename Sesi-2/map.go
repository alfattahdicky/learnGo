package main

import "fmt"

func main()  {
	// map => key:value (pairs) ,termasuk dalam reference type
	// zero value dari tipe data map yaitu nil
	
	//pendahuluan

	// first way
	// var person map[string]string // declare
	// person = map[string]string{} // inisialisasi

	// person["name"] = "Dicky"
	// person["age"] = "18"

	// second way
	// var person = map[string]string {
	// 	"name": "dicky",
	// 	"age": "18",
	// }
	// fmt.Println("name:", person["name"]);
	// fmt.Println("age:", person["age"]);

	// loop map
	// for key, value := range person {
	// 	fmt.Println(key, ":", value)
	// }

	// delete value
	// delete(person, "age")
	// fmt.Println("deleting value", person)

	// detecting value => menggunakan variable penampung
	// delete(person, "age")
	// value, exist := person["age"]

	// if exist {
	// 	fmt.Println(value)
	// }else {
	// 	fmt.Println("value does'nt exist")
	// }

	// combining slice with map
	var people = []map[string]string {
		{"name" : "Dicky", "age": "24"},
		{"name": "Azka", "age": "20"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s \n", i, person["name"], person["age"])
	}
}