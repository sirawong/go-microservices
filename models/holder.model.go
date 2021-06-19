package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HolderModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	First_name string             `bson:"first_name"`
	Last_name  string             `bson:"last_name"`
	Email      string             `bson:"email"`
	Held_books []string           `bson:"held_books"`
}
