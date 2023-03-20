package usermodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDTO struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID        string             `json:"userId,omitempty" bson:"userId,omitempty"`
	DisplayName   string             `json:"displayName,omitempty" bson:"displayName,omitempty"`
	PictureURL    string             `json:"pictureUrl,omitempty" bson:"pictureUrl,omitempty"`
	StatusMessage string             `json:"statusMessage,omitempty" bson:"statusMessage,omitempty"`
	Language      string             `json:"language,omitempty" bson:"language,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
