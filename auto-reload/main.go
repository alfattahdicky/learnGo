package main

import(
	"auto-reload/routers"
)

func main()  {

	var PORT = ":5174"

	routers.StartServer().Run(PORT)
}