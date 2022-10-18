package main

import "fmt"

func main()  {
	// slice tidak memiliki sifat fixed-length => termasuk kategori reference

	// var fruits = make([]string, 3)
	// pembuatan slice 

	// var somethings = []string {"Apple", "banana", "manggo"}

	// _,_ = somethings, fruits

	// fmt.Printf("%#v \n", somethings)

	// menambah elemen
	// fruits[0] = "manggo"
	// fruits[1] = "apple"
	// fruits[2]  = "cheese"

	// fruits = append(fruits, "manggo", "apple", "manggo")

	// fmt.Printf("%#v \n", fruits)

	// gabung 2 array jadi satu => menggunakan ellipsis
	// var fruits1 = []string {"Apple", "banana", "manggo"}
	// var fruits2 = []string {"durian", "pineapple"}

	// fruits1 = append(fruits1, fruits2...)
	// fmt.Printf("%#v \n", fruits1)

	// mengcopy element slice ke slice lainnya arah copy fruits2 ke fruits 1
	// copy(destinasi, source)
	// fruit := copy(fruits1, fruits2)

	// fmt.Println("Fruits 1 =>", fruits1)
	// fmt.Println("Fruits 2 =>", fruits2)
	// fmt.Println("Copy Fruits =>", fruit)

	// mengambil beberapa element (slicing) slice(start:stop) 
	// pengambilan sampai stop - 1 
	// var fruits = []string {"Apple", "banana", "manggo", "pineapple"}
	
	// fruits1 := fruits[1:3]

	// fruits1 := fruits[:] // mendapatakan seluruh element

	// fmt.Println("fruits 1", fruits1)


	// combining slicing and append => replace sesuai index
	// var fruits = []string {"Apple", "banana", "manggo", "pineapple"}

	// fruits = append(fruits[0:3], "rambutan")

	// fmt.Println(fruits)

	// backing array => menyimpan element pada slice terhdap slice yang sudah ada
	// var fruits1 = []string {"Apple", "banana", "manggo"}
	// var fruits2 = fruits1[1:3]

	// fruits2[0] = "pineapple"

	// // hal ini dikarenakan fruits1 dan fruits2 masih satu backing array yang sama
	// fmt.Println("fruit 1 =>", fruits1) // fruit 1 => [Apple pineapple manggo]
	// fmt.Println("fruit 2 =>", fruits2) // fruit 2 => [pineapple manggo]


	// cap function => mengetahui kapasitas dari sebuah array maupun slice 
	// panjang dan kapasitas akan berubah ketika index slice dimulai dari 1
	// var fruits1 = []string {"Apple", "banana", "manggo"}

	// fruit := fruits1[0:2]

	// fmt.Println("Cap function", cap(fruit)); // index : 1 => 2 // index 0:2 => 3
	// fmt.Println("len function", len(fruit)); // index : 1 => 2 // index 0: 2 => 2


	// creating new backing array
	cars := []string{"Ford", "Audi", "Honda", "Rover"}

	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println("cars:", cars)
	fmt.Println("newCars", newCars)

}