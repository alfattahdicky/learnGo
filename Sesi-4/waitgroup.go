package main

import (
	"fmt"
	"sync"
)

func main()  {
	// waitgroup => struct dari package sync untuk sinkronisasi terhadap go routine
	// cara menggunakan function sleep untuk menahan main function bukan cara terbaik
	
	fruits := []string{"apple", "banana", "durian", "manggo"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1) // menambah counter untuk memberitahu waitgroup tentang jumlah go routine yang harus ditunggu
		go printFruit(index, fruit, &wg)
	}

	wg.Wait() // menunggu seluruh go rountine menyesaikan prosessnya untuk menahan func main hingga seluruh proses go routine selesai
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit -> %s \n", index, fruit)
	wg.Done()
}