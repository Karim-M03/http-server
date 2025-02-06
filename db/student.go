package db

import ("go.mongodb.org/mongo-driver/bson/primitive")

type Student struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string   `bson:"name" json:"name"`
	LastName string   `bson:"last_name" json:"last_name"`
	Email	 string   `bson:"email" json:"email"`
	BirthDate primitive.DateTime   `bson:"birth_date" json:"birth_date"`
	Exams	[]primitive.ObjectID	`bson:"exams" json:"exams"`
}