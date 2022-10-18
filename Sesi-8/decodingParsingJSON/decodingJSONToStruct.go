package decodingJSONToStruct

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Age int `json:"age"`
}

func main()  {
	var jsonString = `
		{
			"full_name": "Dicky AL Fattah",
			"email":"alfattahdicky@gmail.com",
			"age": 24
		}
	`

	var result Employee
	// convert decode json to struct using slice of byte
	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result.FullName)
	fmt.Println("email:", result.Email)

}