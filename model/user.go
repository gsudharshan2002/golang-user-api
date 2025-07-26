package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name        string             `json:"name"`
    DOB         string             `json:"dob"`
    Address     string             `json:"address"`
    Description string             `json:"description"`
    CreatedAt   string             `json:"createdAt"`
}
