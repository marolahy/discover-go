package controllers

import "net/http"

// RegisterController is a front controller
func RegisterController() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}
