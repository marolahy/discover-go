package main

import (
	"errors"
	"fmt"

	"github.com/marolahy/discover-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Beza",
		LastName:  "RAND",
	}
	fmt.Println(u)
}

func startWebserver(port int) error {
	fmt.Println("Start webserver")

	return errors.New("Somithing went wrong")
}
