package util

import (
	"bytes"
	"encoding/gob"
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
	w.WriteHeader(statusCode)
	log.WithFields(log.Fields{
		"status_code": statusCode,
	}).Debug("HTTP response sent")
	fmt.Fprintf(w, JSONStringify(res))
}

func ToBytes(key interface{}) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		log.Error(err)
	}
	return buf.Bytes()
}
