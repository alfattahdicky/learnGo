package main

import (
	"auto-reload/routers"
	_ "fmt"
	_ "math/rand"
	_ "time"
)

func main()  {

	var PORT = ":5174"

	routers.StartServer().Run(PORT)

	// fmt.Println(rand.Intn(10))
	// <-time.After(4 * time.Second)
	// fmt.Println("expired")

	

}