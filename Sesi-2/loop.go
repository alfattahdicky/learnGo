package main

import "fmt"

func main()  {
	// first way
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka ", i)
	// }

	// var i = 0
	// second way
	// for i < 5 {
	// 	fmt.Println("Angka", i);
	// 	i++
	// }

	// third way
	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if(i == 5) {
	// 		break
	// 	}
	// }

	// bisa menerapkan break & continue keywords
	// for i := 0; i < 10; i++ {
	// 	if i % 2 != 1 {
	// 		continue
	// 	}

	// 	// if i > 8 {
	// 	// 	break
	// 	// }

	// 	fmt.Println("Angka", i)
	// }

	// nested loop 
	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// Loopings (label) => memberhentikan loop paling luar dengan keyword outerLoop
	outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Println("Perulangan ke - 1", i +1)
		for j := i; j < 5; j++ {
			if i == 2  {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}


}