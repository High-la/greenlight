package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with info about the
// app status, operating envt and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// Create a map which holds the infn that we want to send in the response.
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	// pass the map to the jsno.Marshal() fun. This returns a []byte slice
	// containing the encoded JSON. If there is an error, we log it and send the client
	// a generic error message.
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	// Append a newline to the JSON. This is just a small nicety to make it easier to
	// view in terminal apps.
	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	// Use w.Write to send the []byte slice containing the JSON as response body.
	w.Write(js)
}
