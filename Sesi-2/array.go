package main

import "fmt"

func main()  {
	// array didalam go memiliki sifat fixed length/ panjang yang tetap karena sudah ditentukan diawal
	// var numbers [4] int
	// numbers = [4]int{2, 4,6 ,7}
	// var strings = [3]string{"Dicky", "Azka", "Diaz"}

	// fmt.Printf("%#v \n", numbers)
	// fmt.Printf("%#v \n", strings)

	// modifikasi array
	// var fruits = [5]string {"Mangga", "jeruk", "pisang"}
	// fruits[0] = "Anggur"
	// fruits[1] = "Jambu"
	
	// fmt.Printf("%#v\n", fruits) // [5]string{"Anggur", "Jambu", "pisang", "", ""}

	// looping array element
	// var fruits = [3]string {"Jambu", "Mangga","jeruk"}


	// first way
	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("Index: %d, Value: %s\n", i , fruits[i])
	// }

	// second way
	// for i, v := range fruits {
	// 	fmt.Printf("Index: %d, Value: %s\n", i , v)
	// }

	// multidimensional array
	balances := [2][3]int{{4,5,6}, {6,7,8}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}