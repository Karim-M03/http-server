package db

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Exam struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Lecture     string   `bson:"name" json:"name"`
	LastName string   `bson:"last_name" json:"last_name"`
	BirthDate primitive.DateTime   `bson:"birth_date" json:"birth_date"`
	Exams	[]primitive.ObjectID	`bson:"exams" json:"exams"`
}