package db

import ("go.mongodb.org/mongo-driver/bson/primitive")

type History struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Lecture  string   `bson:"name" json:"name"`
	Student  []Student `bson:"students" json:"students"`
	Semester	string `bson:"semester" json:"semester"`
}