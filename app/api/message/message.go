package message

import (
	"linebot-go/app/models/messagemodel"
	"linebot-go/services/ginservices"
	"linebot-go/services/mongodb"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func GetMessages(c *gin.Context) {
	type structRequest struct {
		UserID string `json:"user_id,omitempty" form:"user_id,omitempty" binding:"-"`
		messagemodel.PageQueryArgs
	}
	var reqJSON structRequest
	_, err := ginservices.GinRequest(c, &reqJSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Create mongo client
	mgClient, err := mongodb.NewMongoClient()
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	userDAO := messagemodel.NewMessageDAO(mgClient)
	skip := (reqJSON.Page - 1) * reqJSON.Limit

	filter := bson.M{}
	if reqJSON.UserID != "" {
		filter = bson.M{
			"userId": reqJSON.UserID,
			// "$and":   and,
		}
	}

	if reqJSON.Keyword != "" {
		and := []bson.M{
			{"messageText": bson.M{"$regex": primitive.Regex{Pattern: reqJSON.Keyword, Options: "i"}}},
		}
		filter["$and"] = and
	}

	result, _ := userDAO.GetByQuery(filter, int64(reqJSON.Limit), int64(skip), reqJSON.Order, reqJSON.By)
	if len(*result) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	total, err := userDAO.CountDocuments(filter)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.Writer.Header().Set("x-page", strconv.Itoa(reqJSON.Page))
	c.Writer.Header().Set("x-per-page", strconv.Itoa(reqJSON.Limit))
	c.Writer.Header().Set("x-total", strconv.Itoa(int(total)))
	totalPages := math.Ceil(float64(total) / float64(reqJSON.Limit))
	c.Writer.Header().Set("x-total-pages", strconv.Itoa(int(totalPages)))
	c.JSON(http.StatusOK, result)
}
