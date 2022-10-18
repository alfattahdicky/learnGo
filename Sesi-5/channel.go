package main

import (
	"fmt"
	"time"
)

func main()  {
	// channel => mekanisme untuk goroutine saling berkomunikasi yaitu pertukaran data antara goroutine satu dengan goroutine lainnya
	// keyword chan
	// c := make(chan string)

	// // tanda <- yaitu operator pengiriman data dari goroutine satu dengan yang lainnya
	// c <- value // mengirim data ke channel
	
	// result := <- c // menerima data dari channel

	// implementing channel

	// c := make(chan string)

	// // mengirim variable result ke channel	
	// go introduce("Dicky", c)
	// go introduce("Azka", c)

	// msg1 := <- c // menerima data dari channel
	// fmt.Println(msg1)

	// msg2 := <- c
	// fmt.Println(msg2)

	// close(c) // channel sudah tidak digunakan untuk berkomunikasi lagi

	// channel with anonymous function
	// c := make(chan string)

	// students := []string {"Dicky", "Azka", "Diaz"}

	// for _, v := range students{
	// 	go func (student string)  {
	// 		fmt.Println("student", student)
	// 		result := fmt.Sprintf("Hai, my name is %s", student)

	// 		c <- result // mengirim variable result ke channel
	// 	}(v)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(c)
	// }

	// close(c)

	// channel direction
	// c := make (chan string)

	// students := []string{"Dicky", "Azka", "dicaz"}

	// for _, v := range students {
	// 	go introduce(v, c)
	// }

	// for i := 1; i <= 3; i++ {
	// 	print(c)
	// }

	// close(c)

	// UnBuffered Channel => pengiriman dan penerimaan data bersifat synchronous
	// apabila ditentukan kapasitasnya dari buffernya pengiriman data akan bersifat async
	// c1 := make(chan int)

	// go func(c chan int)  {
	// 	fmt.Println("func goroutine start send data ")
	// 	c <- 20
	// 	fmt.Println("func goroutine end send data") // tertahan 
	// }(c1)

	// fmt.Println("main goroutine for 2 seconds")
	// time.Sleep(time.Second * 2)

	// fmt.Println("main goroutine start receiving data")
	// d := <- c1
	// fmt.Println("main goroutine end receiving data", d)

	// close(c1)

	// time.Sleep(time.Second)

	// // buffered channel 
	// c1 := make(chan int, 3)

	// go func(c chan int)  {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Println("func start send data", i)
	// 		c <- i
	// 		fmt.Println("func end send data", i)
	// 	}

	// 	close(c)
	// }(c1)

	// fmt.Println("main goroutine sleeps 2 seconds")

	// time.Sleep(time.Second * 2)

	// for v := range c1 {
	// 	fmt.Println("main goroutine receive value", v)
	// }

	// channel select => komunikasi lebih dari satu channel 
	c1 := make(chan string)
	c2 := make(chan string)

	go func()  {
		time.Sleep(2 * time.Second)
		
		c1 <- "Hello!"
	}()

	go func()  {
		time.Sleep(2 * time.Second)
		
		c2 <- "World!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <- c1 :
			fmt.Println("received", msg1)
		case msg2 := <- c2 :
			fmt.Println("Received", msg2)
		}
	}

}

// channel with anonymous function
// func print(c chan string)  {
// 	fmt.Println(<-c) // menerima data 
// }

// implementing channel
// func introduce(student string, c chan string)  {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
	
// 	c <- result // mengirim variable result ke channel
// }

// // channel direction
// // print => menerima data
// func print(c <-chan string)  {
// 	fmt.Println(<-c) // menerima data 
// }

// // introduce => mengirim data
// func introduce(student string, c chan<- string)  {
// 	result := fmt.Sprintf("Hai, my name is %s", student)
	
// 	c <- result // mengirim variable result ke channel
// }

