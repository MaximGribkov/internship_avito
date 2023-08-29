package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Errors struct {
	Massage string `json:"massage"`
}

// Обработчик ошибок
func newErrorResponse(c *gin.Context, statusCode int, massage string) {
	logrus.Error(massage)
	c.AbortWithStatusJSON(statusCode, Errors{Massage: massage})
}
