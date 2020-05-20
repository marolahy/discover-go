package main

import (
	"fmt"
	"net/http"

	"github.com/marolahy/discover-go/controllers"
	"github.com/marolahy/discover-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Beza",
		LastName:  "RAND",
	}
	fmt.Println(u)
	startWebserver(3000)
}

func startWebserver(port int) error {
	fmt.Println("Start webserver")

	controllers.RegisterController()
	http.ListenAndServe(fmt.Sprint(":", port), nil)

	return nil
}
