package main

import (
	"fmt"
	"encoding/json"
)

type Employee struct {
	FullName string `json:"full_name"`
	Age int `json:"age"`
}


func main()  {
	var jsonString = `[
		{
			"full_name": "dicky al fattah",
			"age": 20
		},
		{
			"full_name": "azka faiz",
			"age": 18
		}
	]`

	var result []Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("index %d: %+v \n", i+1, v)
	}
}