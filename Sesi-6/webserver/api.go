// Get & POST

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type Employee struct {
	ID int
	Name string
	Age int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Dicky", Age: 20, Division: "IT"},
	{ID: 2, Name: "Azka", Age: 17, Division: "Security"},
	{ID: 3, Name: "Diaz", Age: 18, Division: "Marketing"},
}

var PORT = ":8080"

func main()  {
	// get employees
	http.HandleFunc("/employees", getEmployees)

	// post employess
	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Application Listening Port", PORT)

	http.ListenAndServe(PORT, nil)
}

// GET Employees
func getEmployees(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		tpl, err := template.ParseFiles("./Sesi-6/template.html")


		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, employees)

		json.NewEncoder(w).Encode(employees)
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

// POST Employees 
func createEmployee(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		division := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)

		if(err != nil) {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		newEmployee := Employee {
			ID: len(employees) + 1,
			Name: name,
			Age: convertAge,
			Division: division,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)

		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}