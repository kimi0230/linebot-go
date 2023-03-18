package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDTO struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        string             `bson:"userId,omitempty"`
	DisplayName   string             `bson:"displayName,omitempty"`
	PictureURL    string             `bson:"pictureUrl,omitempty"`
	StatusMessage string             `bson:"statusMessage,omitempty"`
	Language      string             `bson:"language,omitempty"`
}