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
		user.POST("/create", h.createUser)
		user.POST("/add-slug", h.addUserToSlug)
		user.DELETE("/delete-slug", h.deleteUserFromSlug)
		user.GET("/get-user", h.getUserSlug)

	}

	segments := router.Group("/segments")
	{
		segments.POST("/create", h.createSegments)
		segments.DELETE("/delete", h.deleteSegments)
		segments.GET("/get-slug", h.getSegmentsUser)
	}

	return router
}
