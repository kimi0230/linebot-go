package ginservices

import (
	"bytes"
	"io/ioutil"

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
