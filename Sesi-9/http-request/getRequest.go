package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.Body)
	// mengubah nilai yang diakses menjadi slice of byte
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	// convert slice of byte to string
	sb := string(body)

	log.Println(sb)
}