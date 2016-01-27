package httpx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	return
}

func ReadJson(r *http.Response, v interface{}) (err error) {
	if r.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Status: %s", http.StatusText(r.StatusCode))
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, v)
	return
}

func GetJson(url string, v interface{}) (err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()
	return ReadJson(r, v)
}
