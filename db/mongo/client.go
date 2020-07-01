package mongo

import (
	"context"
	"time"

	"github.com/seanarwa/common/config"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client = nil

func Connect(connectionUrl string) {

	logFields := log.Fields{"connection_url": connectionUrl}

	var err error = nil

	mongoClient, err = mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		log.WithFields(logFields).Fatal("Error occured when trying to initialize MongoDB client: ", err)
	} else {
		log.WithFields(logFields).Debug("Initialized MongoDB client")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.WithFields(logFields).Fatal("Error occured when trying to connect to MongoDB: ", err)
	} else {
		log.WithFields(logFields).Debug("MongoDB client has connected to MongoDB")
	}
}

func Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := mongoClient.Disconnect(ctx)
	if err != nil {
		log.Fatal("Error occured when trying to disconnect from MongoDB: ", err)
	} else {
		log.Debug("MongoDB client has disconnected from MongoDB")
	}
}

func getDefaultDatabase() *mongo.Database {
	databaseName := config.GetString("db.mongo.database_name")
	return mongoClient.Database(databaseName)
}

func getDefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func Test() {
	collection := mongoClient.Database("test").Collection("test")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection.Drop(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	log.Info(id)

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
