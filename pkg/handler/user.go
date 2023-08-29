package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

// Метод создания пользователя
func (h *Handler) createUser(c *gin.Context) {
	var input model.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	answer, err := h.services.LogicsUser.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": answer,
	})
}

// Метод добавления сегмента пользователю
func (h *Handler) addUserToSegment(c *gin.Context) {
	var input model.UserSegments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	answer, err := h.services.LogicsUser.AddUserToSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":      input.Id,
		"segment_name": input.SegmentsName,
		"result":       answer,
	})
}

// Метод удаление сегментов у пользователя
func (h *Handler) deleteUserFromSegment(c *gin.Context) {
	var input model.UserSegments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	answer, err := h.services.LogicsUser.DeleteUserFromSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":      input.Id,
		"segment_name": input.SegmentsName,
		"result":       answer,
	})
}

// Метод получения списка сегментов пользователя
func (h *Handler) getUserSegment(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	answer, err := h.services.LogicsUser.GetUserSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": input.Id,
		"result":  answer,
	})
}
