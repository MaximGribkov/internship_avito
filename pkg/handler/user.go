package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

// @Summary createUser
// @Description Метод создания пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input body model.User true "user"
// @Success      200  {string}  ok
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /user/create [post]
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

// @Summary addUserToSegment
// @Description Метод добавления сегмента пользователю
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input body model.User true "user id and list segment"
// @Success      200  {string}  ok
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /user/add-segment [post]
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

// @Summary deleteUserFromSegment
// @Description Метод удаление сегментов у пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        input body model.User true "user id and list segment"
// @Success      200  {string}  ok
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /delete-segment [delete]
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

// @Summary getUserSegment
// @Description Метод получения списка сегментов пользователя
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path int true "user id"
// @Success      200  {string}  model.User
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /get-user-segment [get]
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
