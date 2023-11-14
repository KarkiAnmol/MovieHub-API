// Package utils provides utility functions for handling HTTP requests in the MovieHub API.

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ParseBody reads the request body and unmarshals the JSON data into the specified interface 'x'.
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// Unmarshal JSON 'body' into interface 'x'.
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
