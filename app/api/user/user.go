package user

import (
	"bytes"
	"io/ioutil"
	"linebot-go/app/models/usermodel"
	"linebot-go/services/mongodb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	result, _ := userDAO.GetByQuery(reqJSON.Keyword, int64(reqJSON.Limit), int64(skip), reqJSON.Order, reqJSON.By)
	// if result == nil {
	// 	c.AbortWithStatus(http.StatusNoContent)
	// }

	c.JSON(http.StatusOK, result)
}
