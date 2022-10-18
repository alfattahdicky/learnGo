package main

import "fmt"

func main()  {
	// var currentYear int = 2022

	// if else (temporary variable)

	// if previousNumber := currentYear - 2001; previousNumber > 17 {
	// 	fmt.Println("Kamu boleh membuat sim");
	// }else {
	// 	fmt.Println("kamu tidak boleh membuat sim");
	// }

	// if currentYear > 2017 {
	// 	fmt.Println("Tahun Modern")
	// }else {
	// 	fmt.Println("Tahun Lama")
	// }

	// switch in go tidak ada keyword break
	// keyword fallthrough yaitu melanjutkan pengecekan ke case selanjutnya meskipun case telah terpenuhi kondisinya
	var score int = 1
	var result string = ""
	switch  {
	case score == 0: result += "Portugal Menang" 
	case (score > 0) && (score < 2): 
		result += "Inggris Menang"
		fallthrough
	default: 
		{
			result += "tidak ada yang menang" 
		}
	}

	fmt.Println(result)

}