package main

import (
	"fmt"
	"os"
)

func main()  {
	// defer => memanggil sebuah function yang dimaan proses eksekusi akan ditahan hingga block function selesai

	// defer fmt.Println("defer function execute") // last execution
	// fmt.Println("hai")
	// fmt.Println("welcome back to golang")

	// // defer #2 
	// callDeferFunc()
	// fmt.Println("haii")

	// exit => menghentikan suatu program secara paksa
	// defer fmt.Println("Invoke with defer")  // tidak ditampilkan
	// fmt.Println("Before exiting")
	// os.Exit(1)
}

// func callDeferFunc() {
// 	defer deferFunc() // terjadi bukan didalam blok main
// }

// func deferFunc() {
// 	fmt.Println("Defer func execution") // first print 
// }