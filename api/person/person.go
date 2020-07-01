package person

import (
	"io/ioutil"
	"net/http"

	"github.com/seanarwa/ds/db/mongo"
	"github.com/seanarwa/ds/util"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func Handler(w http.ResponseWriter, req *http.Request) {

	method := req.Method
	path := req.URL.EscapedPath()

	log.Debug(method, " ", path)

	switch method {
	case "GET":
		get(w, req)
		break
	case "POST":
		post(w, req)
		break
	case "PUT":
		put(w, req)
		break
	case "DELETE":
		delete(w, req)
		break
	default:
		log.WithFields(log.Fields{
			"path":   req.URL.EscapedPath(),
			"method": method,
		}).Warn("Unsupported HTTP method has been called")
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	util.WriteHTTPResponse(w, mongo.GetAll("person"), http.StatusOK)
}

func post(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var query bson.D
	err := bson.UnmarshalExtJSON(body, false, &query)
	if err != nil {
		log.Error(err)
	}
	insertedID := mongo.InsertOne("person", query)
	util.WriteHTTPResponse(w, map[string]interface{}{
		"inserted_id": insertedID,
	}, http.StatusOK)
}

func put(w http.ResponseWriter, req *http.Request) {

}

func delete(w http.ResponseWriter, req *http.Request) {

}
