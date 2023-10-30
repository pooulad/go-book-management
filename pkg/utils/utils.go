package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParsBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}
