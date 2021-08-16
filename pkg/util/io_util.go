package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ParseResponseWithJSON will parse the response's body.
// Try to unmarshal by json format.
func ParseResponseWithJSON(response *http.Response, obj interface{}) error {
	if response == nil {
		return fmt.Errorf("response is nil")
	}
	if response.Body == nil {
		return fmt.Errorf("response-body is nil, can't read")
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, obj)
	return err
}