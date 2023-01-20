package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID    primitive.ObjectID
	Title string
	Name  string
}
