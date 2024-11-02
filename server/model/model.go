package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Blog struct{
    Id bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Title string `json:"title,omitempty" bson:"title,omitempty"`
    Body string `json:"body,omitempty" bson:",omitempty"`
}
