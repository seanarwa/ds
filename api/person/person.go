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
		getMethod(w, req)
		break
	case "POST":
		postMethod(w, req)
		break
	case "PUT":
		putMethod(w, req)
		break
	case "DELETE":
		deleteMethod(w, req)
		break
	default:
		log.WithFields(log.Fields{
			"path":   req.URL.EscapedPath(),
			"method": method,
		}).Warn("Unsupported HTTP method has been called")
	}
}

func getMethod(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var query bson.D
	err := bson.UnmarshalExtJSON(body, false, &query)
	if err != nil {
		log.Error(err)
	}
	id := query.Map()["_id"]
	if id == nil {
		util.WriteHTTPResponse(w, mongo.GetAll("person"), http.StatusOK)
	} else {
		util.WriteHTTPResponse(w, mongo.FindOne("person", id.(string)), http.StatusOK)
	}
}

func postMethod(w http.ResponseWriter, req *http.Request) {
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

func putMethod(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var query bson.D
	err := bson.UnmarshalExtJSON(body, false, &query)
	if err != nil {
		log.Error(err)
	}
	id := query.Map()["_id"]
	if id == nil {
		log.Error("Error occured when trying to update one from MongoDB database: _id does not exist in body")
		util.WriteHTTPResponse(w, map[string]interface{}{
			"error": "_id does not exist in body",
		}, http.StatusBadRequest)
		return
	}
	for k, v := range query {
		if v.Key == "_id" {
			query = append(query[:k], query[k+1:]...)
			break
		}
	}
	result := mongo.UpdateOne("person", id.(string), query)
	util.WriteHTTPResponse(w, map[string]interface{}{
		"updated_id":   id.(string),
		"update_count": result.ModifiedCount,
	}, http.StatusOK)
}

func deleteMethod(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	var query bson.D
	err := bson.UnmarshalExtJSON(body, false, &query)
	if err != nil {
		log.Error(err)
	}
	id := query.Map()["_id"]
	if id == nil {
		log.Error("Error occured when trying to delete one from MongoDB database: _id does not exist in body")
		util.WriteHTTPResponse(w, map[string]interface{}{
			"error": "_id does not exist in body",
		}, http.StatusBadRequest)
		return
	}
	result := mongo.DeleteOne("person", id.(string))
	util.WriteHTTPResponse(w, map[string]interface{}{
		"deleted_id":   id,
		"delete_count": result.DeletedCount,
	}, http.StatusOK)
}
