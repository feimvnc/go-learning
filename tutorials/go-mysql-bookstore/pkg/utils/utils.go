package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// pass an empty interfce, any type will satisfy it
// generic functino that can receive any value
// allow to pass either value or pointer references
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
