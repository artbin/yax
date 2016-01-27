package httpx

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	b, err := json.Marshal(error)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	return
}
