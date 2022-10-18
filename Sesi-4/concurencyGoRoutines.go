package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	// concurrency => siapa yang akan menyelesaikan tugasnya terlebih dahulu
	// goroutine => thread ringan untuk melakukan concurrency bersifat asynchronous
	// menggunakan keyword go

	// go goroutine();

	// async #1
	fmt.Println("main execution started")

	go firstProcess(8)
	secondProcess(8)
	fmt.Println("No. of Gorountines", runtime.NumGoroutine()) // mengecek package runtime jumlah gorountine berjalan

	time.Sleep(time.Second * 2) // menahan function main dalam 2 detik
	fmt.Println("main execution ended")
	
}

func firstProcess(index int)  {
	fmt.Println("First process func started")
	for i:= 1; i <= index; i++ {
		fmt.Println("First Process i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int)  {
	fmt.Println("Second process func started")
	for i:= 1; i <= index; i++ {
		fmt.Println("Second Process i=", i)
	}
	fmt.Println("Second process func ended")
}

// func goroutine() {
// 	fmt.Println("Hello")
// }