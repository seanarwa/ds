package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func JSONStringify(obj interface{}) string {
	jsonString, err := json.Marshal(obj)
	if err != nil {
		log.WithFields(log.Fields{
			"object": obj,
		}).Fatal("Error occured when trying to stringify object to json string")
	}
	return string(jsonString)
}

func WriteHTTPResponse(w http.ResponseWriter, res interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseString := JSONStringify(res)
	fmt.Fprintf(w, responseString)
	log.WithFields(log.Fields{
		"status_code": statusCode,
	}).Debug("HTTP response sent")
	log.WithFields(log.Fields{
		"response_body": responseString,
	}).Trace()
}
