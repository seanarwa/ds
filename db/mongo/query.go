package mongo

import (
	"context"

	"github.com/seanarwa/common/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection string, query bson.D) *mongo.Cursor {
	databaseName := config.GetString("db.mongo.database_name")
	cur, err := mongoClient.Database(databaseName).Collection(collection).Find(getDefaultContext(), query)
	logFields := log.Fields{
		"database_name": databaseName,
		"collection":    collection,
	}
	if err != nil {
		log.WithFields(logFields).Error("Error occured when trying to find in MongoDB database: ", err)
	}
	log.WithFields(logFields).Trace("Good find query from MongoDB")
	return cur
}

func GetAll(collection string) interface{} {
	databaseName := config.GetString("db.mongo.database_name")
	logFields := log.Fields{
		"database_name": databaseName,
		"collection":    collection,
	}

	cur := Find(collection, bson.D{})
	var result []bson.M
	err := cur.All(context.Background(), &result)
	if err != nil {
		log.WithFields(logFields).Error("Error occured when trying to get all from MongoDB database: ", err)
	}
	log.WithFields(logFields).Trace("Good get all query from MongoDB")
	return result
}

func InsertOne(collection string, query bson.D) string {
	databaseName := config.GetString("db.mongo.database_name")
	logFields := log.Fields{
		"database_name": databaseName,
		"collection":    collection,
	}
	result, err := mongoClient.Database(databaseName).Collection(collection).InsertOne(getDefaultContext(), query)
	if err != nil {
		log.WithFields(logFields).Error("Error occured when trying to insert into MongoDB database: ", err)
	}
	log.WithFields(logFields).Trace("Good insertion into MongoDB")
	return result.InsertedID.(string)
}
