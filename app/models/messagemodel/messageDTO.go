package messagemodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDTO struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Type        string             `bson:"type,omitempty"`
	UserID      string             `bson:"userId,omitempty"`
	MessageID   string             `bson:"messageId,omitempty"`
	MessageType string             `bson:"messageType,omitempty"`
	MessageText string             `bson:"messageText,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}
