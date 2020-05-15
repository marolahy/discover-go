package main

import (
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
