package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Brave-man/base/bootstrap/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetAllParams 获取请求中的所有参数
func GetAllParams(c *gin.Context) (headers, params map[string]interface{}, err error) {
	headers = map[string]interface{}{}
	params = map[string]interface{}{}
	// 处理GET请求
	if c.Request.Method == "GET" {
		queryData := c.Request.URL.Query()
		for key, value := range queryData {
			params[key] = value[0]
		}
		return
	}

	var contentType string
	contentTypeList, exist := c.Request.Header["Content-Type"]
	if !exist {
		err = errors.New("headers do not carry Content-Type")
		return
	}
	contentType = contentTypeList[0]
	headers["Content-Type"] = contentType

	// 处理非GET请求, 本项目默认非GET请求, 参数只读取postForm和body
	if contentType == "application/json" {
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		err = json.Unmarshal(body, &params)
	} else if contentType == "application/x-www-form-urlencoded" {
		if err = c.Request.ParseForm(); err == nil {
			formData := c.Request.PostForm
			for key, value := range formData {
				params[key] = value[0]
			}
		}
	}
	return
}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.Log

		// params 请求参数
		headers, params, err := GetAllParams(c)
		// headers 中的Content-Type
		contentType := headers["Content-Type"]
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 0,
				"msg":  fmt.Sprintf("%v", err),
			})

			log.Info("Bad Request",
				zap.String("requestMethod", c.Request.Method),
				zap.String("requestURI", c.Request.RequestURI),
				zap.Any("Content-Type", contentType),
				zap.String("requestProto", c.Request.Proto),
				zap.String("requestUseragent", c.Request.UserAgent()),
				zap.String("requestReferer", c.Request.Referer()),
				zap.Any("requestData", params),
				zap.String("requestClientIP", c.ClientIP()),
				zap.Int("responseStatusCode", c.Writer.Status()),
				zap.String("errorMessage", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			)
			c.Abort()
		}
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()

		log.Info("Request",
			zap.String("requestMethod", c.Request.Method),
			zap.String("requestURI", c.Request.RequestURI),
			zap.Any("Content-Type", contentType),
			zap.String("requestProto", c.Request.Proto),
			zap.String("requestUseragent", c.Request.UserAgent()),
			zap.String("requestReferer", c.Request.Referer()),
			zap.Any("requestData", params),
			zap.String("latencyTime", fmt.Sprintf("%v", endTime.Sub(startTime))),
			zap.String("requestClientIP", c.ClientIP()),
			zap.Int("responseStatusCode", c.Writer.Status()),
			zap.String("errorMessage", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		)
	}
}
