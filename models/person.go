package models

type Person struct {
	UserID string `bson:"userid" json:"userid"`
	Name   string `bson:"name" json:"name"`
}
