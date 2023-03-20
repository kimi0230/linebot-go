package user

import (
	"bytes"
	"io/ioutil"
	"linebot-go/app/models/usermodel"
	"linebot-go/services/mongodb"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func GinRequest(c *gin.Context, reqJSON interface{}) (interface{}, error) {
	var reqData interface{}
	if c.Request.Method == "GET" {
		reqData = c.Request.URL.Query()
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}
	} else {
		data, err := c.GetRawData()
		if err != nil {
			return nil, err
		}
		reqData = ioutil.NopCloser(bytes.NewBuffer(data))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		if bindErr := c.ShouldBind(reqJSON); bindErr != nil {
			return nil, bindErr
		}

	}
	return reqData, nil
}

func GetUsers(c *gin.Context) {
	type structRequest struct {
		usermodel.PageQueryArgs
	}
	var reqJSON structRequest
	_, err := GinRequest(c, &reqJSON)
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
	userDAO := usermodel.NewUserDAO(mgClient)
	skip := (reqJSON.Page - 1) * reqJSON.Limit

	filter := bson.M{
		"$or": []bson.M{
			// "i" 表示忽略大小寫
			{"displayName": bson.M{"$regex": primitive.Regex{Pattern: reqJSON.Keyword, Options: "i"}}},
		},
	}

	result, _ := userDAO.GetByQuery(filter, int64(reqJSON.Limit), int64(skip), reqJSON.Order, reqJSON.By)
	if len(*result) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
	}
	total, err := userDAO.CountDocuments(filter)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	// https://docs.gitlab.com/ee/api/rest/#other-pagination-headers
	// x-next-page	The index of the next page.
	// x-page	The index of the current page (starting at 1).
	// x-per-page	The number of items per page.
	// x-prev-page	The index of the previous page.
	// x-total	The total number of items.
	// x-total-pages	The total number of pages.

	c.Writer.Header().Set("x-page", strconv.Itoa(reqJSON.Page))
	c.Writer.Header().Set("x-per-page", strconv.Itoa(reqJSON.Limit))
	c.Writer.Header().Set("x-total", strconv.Itoa(int(total)))
	totalPages := math.Ceil(float64(total) / float64(reqJSON.Limit))
	c.Writer.Header().Set("x-total-pages", strconv.Itoa(int(totalPages)))
	c.JSON(http.StatusOK, result)
}