package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Author string             `bson:"author"`
	Title  string             `bson:"title"`
}
