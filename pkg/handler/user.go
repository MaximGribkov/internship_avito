package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.LogicsUser.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) addUserToSlug(c *gin.Context) {

}

func (h *Handler) deleteUserFromSlug(c *gin.Context) {

}

// количество пользователей в сегменте
func (h *Handler) getUserSlug(c *gin.Context) {

}
