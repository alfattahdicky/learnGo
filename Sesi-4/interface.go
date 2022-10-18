package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi + c.radius
}

func (c circle) volume() float64 {
	return 4/3 * math.Pi * math.Pow(c.radius, 3)
}

func main()  {
	// interface => tipe data go yaitu kumpulan dari satu atau lebih method
	// biasanya digunakan untuk implementasi method 
	
	var c1 shape = circle{radius: 5}

	// fmt.Printf("%T \n", c1)
	// fmt.Println("Area C1", c1.area()) // call method area
	// fmt.Println("Area C2", c1.perimeter()) // call method perimeter
	
	// teknik type assertion => mengembalikan suatu tipe data interface kepada tipe data aslinya
	// c1.(circle).volume()

	value, ok := c1.(circle)

	if ok == true {
		fmt.Printf("Circle value %+v \n", value)
		fmt.Printf("Circle volume: %f \n", value.volume())
	}
}