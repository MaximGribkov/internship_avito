package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.GET("/get-user-segment", h.getUserSegment)
		user.POST("/create", h.createUser)
		user.POST("/add-segment", h.addUserToSegment)
		user.DELETE("/delete-segment", h.deleteUserFromSegment)
	}

	segments := router.Group("/segments")
	{
		segments.GET("/get-count-segment", h.userCountInSegment)
		segments.POST("/create", h.createSegments)
		segments.DELETE("/delete", h.deleteSegments)
	}

	return router
}
