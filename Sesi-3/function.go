package main

import (
	"fmt"
	"strings"
)

func main()  {
	// function 
	// greet("Dicky", "20")

	// return
	// fmt.Println(greet("Hello", []string{"dicky", "azka"}))

	// multiple return
	// var area, circumference = calculate(15)
	// fmt.Println("area", area)
	// fmt.Println("circumference", circumference)

	// predefined return value
	// var area, circumference = calculate(15)

	// fmt.Println("area", area)
	// fmt.Println("circumference", circumference)

	// variadic function 
	// fmt.Printf("%v \n", print("dicky", "azka"))

	// variadic function v2
	// fmt.Println("Result", sum([]int{1,2,3,4,5}...))
	// variadic function v3
	profile("Dicky", "pasta", "manggo", "mie")
}

// pengenalan function
// func greet(name, age string)  {
// 	fmt.Println("Hello, my name is", name)
// 	fmt.Println("My age is", age)
// }

// using return 
// func greet(msg string, names []string) string {
// 	var joinStr = strings.Join(names, " ")
// 	var result string = fmt.Sprintf("%s %s", msg, joinStr) // Sprintf mreturn sebuah nilai

// 	return result
// }

// return multiple values
// func calculate(d float64) (float64, float64)  {
// 	var area float64 = math.Pi * math.Pow(d/2, 2)
// 	var circumference = math.Pi * d

// 	return area, circumference
// }

// predefined return value
// return value memberikan sebuah variable sebagai hasil return
// diperlukan mereassign variable yang digunakan dari nilai return
// func calculate(d float64) (area float64, circumference float64) {
// 	area = math.Pi * math.Pow(d/2, 2)
// 	circumference = math.Pi * d

// 	return
// }

// variadic function 
// function yang menerima argument yang tidak terbatas jumlahnya
// func print(names ...string) []map[string]string {
// 	var result []map[string]string
// 	for i, v := range names {
// 		key := fmt.Sprintf("student%d", i + 1)
// 		temp := map[string]string {
// 			key: v,
// 		}
// 		result = append(result, temp)
// 	}

// 	return result
// }

// variadic function v2
// func sum(numbers ...int) int  {
// 	total := 0
// 	for _, v := range numbers {
// 		total += v
// 	}

// 	return total
// }

// variadic function v3
func profile(name string, favFoods ...string)  {
	mergeFavFoods := strings.Join(favFoods, ",")
	fmt.Println("Hello I'm", name)
	fmt.Println("Iam really love", mergeFavFoods)
}



