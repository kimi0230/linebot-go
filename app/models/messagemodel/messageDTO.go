package messagemodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageDTO struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	UserID      string             `json:"userId,omitempty" bson:"userId,omitempty"`
	ReplyToken  string             `json:"replyToken,omitempty" bson:"replyToken,omitempty"`
	MessageID   string             `json:"messageId,omitempty" bson:"messageId,omitempty"`
	MessageText string             `json:"messageText,omitempty" bson:"messageText,omitempty"`
	Timestamp   time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
