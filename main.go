package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Variable creation :=  vs attribution =
	type a int // type a = int

	type Car struct {
		Name  string
		Model string
		Year  int
	}

	func(c Car) Move()
	{
		fmt.Println("The car, " + c.Name + " is moving...")
	}

	func(c Car) Stop()
	{
		fmt.Println("The car, " + c.Name + " stopped...")
	}

	main()
	{
		car1 := Car{Name: "Ford", Model: "Mustang", Year: 1969}
		car2 := Car{Name: "Chevrolet", Model: "Camaro: 1969"}

		// Simple http server
		print("Hello, World\n")
		print("Starting web server...")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			err := json.NewEncoder(w).Encode(car1)
			if err != nil {
				return
			}
	}
		http.ListenAndServe(":8080", nil)
	}


}
