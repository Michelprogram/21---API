package handlers

import (
	"fmt"
	"net/http"
)

func TestRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func ErrorRoute() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var response string = "This API provide GET method for /redis & /notredis."

		fmt.Fprintf(w, "%s", response)
	}
}
