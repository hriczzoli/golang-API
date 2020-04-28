package main

import (
	"net/http"

	"github.com/hriczzoli/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
