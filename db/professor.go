package db



type Professor struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
}