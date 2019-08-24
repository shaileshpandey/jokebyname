package helper

import (
	"encoding/json"
	"net/http"
	"time"
)

// JSONToStruct will return error if unable to serialize json to object
func JSONToStruct(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	defer myClient.CloseIdleConnections()
	return json.NewDecoder(r.Body).Decode(target)
}
