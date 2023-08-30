package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

// @Summary createSegments
// @Description Метод создания сегмента
// @Tags         segment
// @Accept       json
// @Produce      json
// @Param        input body model.Segments true "segment name and(or) percent user"
// @Success      200  {string}  ok
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /segments/create [post]
func (h *Handler) createSegments(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.CreateSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segment_name": answer,
	})
}

// @Summary deleteSegments
// @Description Метод удаления сегмента
// @Tags         segment
// @Accept       json
// @Produce      json
// @Param        input body model.Segments true "segment name"
// @Success      200  {string}  ok
// @Failure      400  {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /segments/delete [delete]
func (h *Handler) deleteSegments(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.DeleteSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": answer,
	})
}

// @Summary userCountInSegment
// @Description Метод получения количества пользователей для одного сегмента
// @Tags         segment
// @Accept       json
// @Produce      json
// @Param        input path int true "segment name"
// @Success      200  {string}  ok
// @Failure      400 {object}  Errors
// @Failure      404 {object}  Errors
// @Failure      500  {object}  Errors
// @Failure      default  {object}  Errors
// @Router       /segments/get-count-segment [get]
func (h *Handler) userCountInSegment(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.UserCountInSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"count": answer,
	})
}
