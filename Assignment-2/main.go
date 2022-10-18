package main

import (
	"assignment-two/database"
	"assignment-two/routers"
)

func main()  {
	database.StartDB()
	db := database.GetDb()
	
	routers.StartServer(db).Run()
}