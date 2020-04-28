package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

//RegisterControllers - function for registering controllers of our webservice and route the incoming requests
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
