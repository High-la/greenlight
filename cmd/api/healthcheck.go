package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with info about the
// app status, operating envt and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
