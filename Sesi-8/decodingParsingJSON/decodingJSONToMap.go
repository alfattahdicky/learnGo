package decodingJSONToMap

import (
	"encoding/json"
	"fmt"
)

func main()  {
	var jsonString = `
		{
			"full_name": "dicky al fattah",
			"age": 24
		}
	`
	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("fullName", result["full_name"])
	fmt.Println("age", result["age"])

}