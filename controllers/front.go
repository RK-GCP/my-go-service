package controllers

import (
	"net/http"
)

//RegisterControllers - register controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
