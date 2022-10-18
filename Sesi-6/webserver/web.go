package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8080"

func main()  {
	http.HandleFunc("/", greet) // routing

	http.ListenAndServe(PORT, nil) // listen port
}

// param 1 => mengirim response balik kepada client yang mengirimkan request
// param 2 => struct mendapatkan semacam url parameter, header dll

func greet(w http.ResponseWriter, r *http.Request)  {
	msg := "Hello World"

	fmt.Fprint(w, msg)
}