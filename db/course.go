package db

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Course struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string   `bson:"name" json:"name"`
	Professor primitive.ObjectID `bson:"professor" json:"professor"`
	Details		string
}