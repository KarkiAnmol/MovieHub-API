package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		//unmarshal JSON 'body' into interface 'x'
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
