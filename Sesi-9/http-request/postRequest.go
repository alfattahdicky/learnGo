package main

import (
	"bytes"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
)

func main()  {
	data := map[string]interface{}{
		"title": "Dicky",
		"body": "lorem ipsum dolor amet",
		"userId": 1,
	}

	// convert map to json
	requestJson, err := json.Marshal(data)

	client := &http.Client{}

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	// eksekusi request
	res, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

}