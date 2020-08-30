package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers - register controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

//encodeResponseAsJSON - this is a helper method encode API to json
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
