package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func (c Car) Move() {
	fmt.Println("The car, " + c.Name + " is moving...")
}

func (c Car) Stop() {
	fmt.Println("The car, " + c.Name + " stopped...")
}

func main() {
	car1 := Car{Name: "Ford", Model: "Mustang", Year: 1969}
	car2 := Car{Name: "Chevrolet", Model: "Camaro", Year: 1969}

	// Simple http server
	//print("Hello, World\n")
	//print("Starting web server...")
	//
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Println("Error starting server:", err)
	//	return
	//}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(car1)
		if err != nil {
			fmt.Println("Error encoding car:", err) // Added error handling
			return
		}
	})

	car1.Move()
	car2.Stop()
}
