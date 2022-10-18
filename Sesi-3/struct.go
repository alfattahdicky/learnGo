package main

import "fmt"

// struct => tipe data yang berupa kumpulan property/field maupun method

// type Employee struct {
// 	name string
// 	age string
// 	division string
// }

// Embedded Struct

type Person struct {
	name string
	age int
}

type Employee struct {
	division string
	person Person
}

func main()  {
	// giving value to struct
	// var employee Employee

	// employee.name = "Dicky"
	// employee.age = "20"
	// employee.division = "Frontend Developer"

	// fmt.Println(employee.name)
	// fmt.Println(employee.age)

	// intializing struct
	// var employee1 = Employee{name: "Azka", age: "17", division: "TNI"}
	// fmt.Printf("Employee => %+v\n", employee1)

	// pointer to a struct
	// var employee1 = Employee{name: "Dicky", age: "20", division: "student"}
	// var employee2 *Employee = &employee1

	// fmt.Println("Employee1", employee1.name)
	// fmt.Println("Employee2", employee2.name)

	// employee1.name = "Azka"

	// fmt.Println("Employee1", employee1.name)
	// fmt.Println("Employee2", employee2.name)

	// embedded struct
	// var employee  = Employee{}

	// employee.person.name = "Dicky"
	// employee.person.age = 20
	// employee.division = "Student"

	// fmt.Printf("%#v \n", employee)

	// anonymous struct
	// // tanpa pengisian field
	// var employee1 = struct {
	// 	person Person
	// 	division string
	// }{}
	// employee1.person.age = 20
	// employee1.person.name = "Dicky"
	// employee1.division = "Frontend Developer"
	// // pengisian field
	// var employee2 = struct {
	// 	person Person
	// 	division string
	// }{
	// 	person: Person{name: "Azka", age: 20},
	// 	division: "TNI",
	// }

	// fmt.Printf("%+v \n", employee1)
	// fmt.Printf("%+v \n", employee2)

	// slice of struct
	var people = []Person {
		{name: "Dicky", age: 20},
		{name: "Azka", age: 15},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	// slice of anonymous struct
	// var employee = []struct {
	// 	person Person
	// 	division string
	// }{
	// 	{person: Person{name: "dicky", age: 20}, division: "Frontend Developer"},
	// 	{person: Person{name: "Azka", age: 15}, division: "TNI"},
	// }

	// for _, v := range employee {
	// 	fmt.Printf("%+v\n", v)
	// }

}